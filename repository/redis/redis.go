package redis

import (
	"TaskQueue/config"
	"TaskQueue/pkg/queue"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(cfg config.Redis) *RedisQueue {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), // آدرس Redis  "localhost:6379"
		Password: cfg.Password,                             // پسورد در صورت نیاز
		DB:       cfg.Db,                                   // شماره دیتابیس (از 0 تا 15 به‌طور پیش‌فرض)
	})

	// Test the connection
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("❌ Failed to connect to Redis: %v\n", err)
		return nil
	}

	fmt.Println("✅ Redis connected successfully")

	return &RedisQueue{
		client: rdb,
	}
}

func (r *RedisQueue) Enqueue(ctx context.Context, queueName string, job *queue.Job) error {
	data, err := json.Marshal(job)
	// fmt.Println(data)
	if err != nil {
		return err
	}
	return r.client.LPush(ctx, "queue:"+queueName, data).Err()
}

func (r *RedisQueue) Dequeue(ctx context.Context, queueName string, timeout time.Duration) (*queue.Job, error) {
	res, err := r.client.BRPop(ctx, timeout, "queue:"+queueName).Result()
	// fmt.Println(res)
	if err != nil {
		return nil, err
	}

	if len(res) != 2 {
		return nil, nil
	}

	var job queue.Job
	if err := json.Unmarshal([]byte(res[1]), &job); err != nil {
		return nil, err
	}

	return &job, nil
}
