package redis

import (
	"TaskQueu/config"
	"TaskQueu/pkg/queue"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
    client *redis.Client
}

func NewRedisQueue(cfg config.Redis) *RedisQueue {
    rdb := redis.NewClient(&redis.Options{
        Addr:   fmt.Sprintf("%s:%s",cfg.Host,cfg.Port), // آدرس Redis  "localhost:6379"
        Password: cfg.Password,               // پسورد در صورت نیاز
        DB:       cfg.Db,                // شماره دیتابیس (از 0 تا 15 به‌طور پیش‌فرض)
    })

    // بررسی اتصال
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Redis connection failed: %v", err)
    }
    fmt.Println("redis connected", pong)

    return &RedisQueue{
        client: rdb,
    }
}

var ctx = context.Background()

func (r *RedisQueue) Enqueue(queueName string, job *queue.Job) error {
    data, err := json.Marshal(job)
    if err != nil {
        return err
    }
    return r.client.LPush(ctx, "queue:"+queueName, data).Err()
}

func (r *RedisQueue) Dequeue(queueName string, timeout time.Duration) (*queue.Job, error) {
    res, err := r.client.BRPop(ctx, timeout, "queue:"+queueName).Result()
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
