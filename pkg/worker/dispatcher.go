package worker

import (
	"TaskQueue/pkg/queue"
	"TaskQueue/repository/postgres"
	"TaskQueue/repository/redis"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	rds "github.com/redis/go-redis/v9"
)

type Dispatcher struct {
	Postgres    *postgres.PostgresDB
	Redis       *redis.RedisQueue  // کلاینت Redis
	QueueName   string             // مثلا "email"
	WorkerCount int                // تعداد worker ها
	JobChan     chan *queue.Job    // کانال ارسال job به worker ها
	WaitGroup   *sync.WaitGroup    // برای مدیریت پایان گوروتین‌ها
	Ctx         context.Context    // برای کنترل context و توقف graceful
	Cancel      context.CancelFunc // لغو context
}

func NewDispatcher(Postgres *postgres.PostgresDB, redis *redis.RedisQueue, queueName string, workerCount int) *Dispatcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &Dispatcher{
		Postgres:    Postgres,
		Redis:       redis,
		QueueName:   queueName,
		WorkerCount: workerCount,
		JobChan:     make(chan *queue.Job),
		WaitGroup:   &sync.WaitGroup{},
		Ctx:         ctx,
		Cancel:      cancel,
	}
}

func (d *Dispatcher) Start() {
	// اجرای Workerها
	for i := 0; i < d.WorkerCount; i++ {
		d.WaitGroup.Add(1)
		go d.startWorker(i)
	}

	go d.dispatchLoop()
}

func (d *Dispatcher) Stop() {
	d.Cancel()
	close(d.JobChan)
	d.WaitGroup.Wait()
}

func (d *Dispatcher) dispatchLoop() {
	fmt.Println("[dispatcher] started listening on queue:", d.QueueName)

	for {
		select {
		case <-d.Ctx.Done():
			fmt.Println("[dispatcher] context canceled, stopping...")
			return

		default:

			job, err := d.Redis.Dequeue(d.Ctx, d.QueueName, 5*time.Second)
			if err != nil {
				if err == rds.Nil {
					continue
				}
				log.Printf("Error BRPOP: %v\n", err)
				continue
			}

			if job == nil {
				continue
			}

			// بررسی زمان اجرای job
			if time.Now().Before(job.RunAt) {
				// Create a copy of the job to avoid race conditions
				jobCopy := *job
				go func(j *queue.Job) {
					select {
					case <-time.After(j.RunAt.Sub(time.Now())):
						select {
						case d.JobChan <- j:
						case <-d.Ctx.Done():
							return
						}
					case <-d.Ctx.Done():
						return
					}
				}(&jobCopy)
			} else {
				d.JobChan <- job
			}
		}
	}
}

func (d *Dispatcher) startWorker(id int) {
	defer d.WaitGroup.Done()

	for job := range d.JobChan {
		log.Printf("[worker-%d] processing job: %s", id, job.ID)

		err := Execute(d.Ctx, job)
		if err != nil {
			log.Printf("[worker-%d] job %s failed: %v", id, job.ID, err)

			job.RetryCount++
			if job.RetryCount <= job.MaxRetries {
				// Exponential backoff: 2^retry_count seconds delay
				backoffDelay := time.Duration(1<<job.RetryCount) * time.Second
				log.Printf("[worker-%d] retrying job %s (%d/%d) after %v delay", id, job.ID, job.RetryCount, job.MaxRetries, backoffDelay)

				// Schedule retry with delay
				jobCopy := *job
				go func(j *queue.Job, delay time.Duration) {
					select {
					case <-time.After(delay):
						select {
						case d.JobChan <- j:
						case <-d.Ctx.Done():
							return
						}
					case <-d.Ctx.Done():
						return
					}
				}(&jobCopy, backoffDelay)
			} else {
				log.Printf("[worker-%d] max retries reached for job %s", id, job.ID)
				// fail log
				d.Postgres.InsertJobLog(postgres.JobLog{
					ID:         uuid.New().String(),
					JobID:      job.ID,
					Queue:      job.Queue,
					Status:     "failed",
					Payload:    job.Payload,
					RetryCount: job.RetryCount,
					Error:      err.Error(),
					CreatedAt:  time.Now(),
				})
			}
		} else {
			log.Printf("[worker-%d] job %s completed successfully", id, job.ID)
			//success log
			d.Postgres.InsertJobLog(postgres.JobLog{
				ID:         uuid.New().String(),
				JobID:      job.ID,
				Queue:      job.Queue,
				Status:     "success",
				Payload:    job.Payload,
				RetryCount: job.RetryCount,
				Error:      "",
				CreatedAt:  time.Now(),
			})
		}
	}
}
