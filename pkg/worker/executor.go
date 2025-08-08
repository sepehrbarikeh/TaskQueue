package worker

import (
	"TaskQueu/pkg/queue"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// تابع اصلی اجرای job
func Execute(ctx context.Context, job *queue.Job) error {
	switch job.Queue {
	case "send_email":
		return handleSendEmail(ctx, job)
	case "process_image":
		return handleProcessImage(ctx, job)
	case "write_log":
		return handleWriteLog(ctx, job)
	default:
		log.Printf("job type %s is unknown", job.Type)
		return errors.New("unknown job type")
	}
}

// =======================
// توابع مربوط به job ها
// =======================

func handleSendEmail(ctx context.Context, job *queue.Job) error {
	rand.Seed(time.Now().UnixNano())

	log.Printf("sending email with payload: %s", job.Payload)

	delay := time.Duration(rand.Intn(2)) * time.Second
	select {
	case <-time.After(delay):
		log.Println("email sent successfully after delay")
		return nil

	case <-ctx.Done():
		return fmt.Errorf("job cancelled during send_email: %w", ctx.Err())
	}
}


func handleProcessImage(ctx context.Context, job *queue.Job) error {
	rand.Seed(time.Now().UnixNano())

	log.Printf("processing image with payload: %s", job.Payload)

	delay := time.Duration(rand.Intn(10)) * time.Second
	select {
	case <-time.After(delay):
		log.Println("process image successfully after delay")
		return nil

	case <-ctx.Done():
		return fmt.Errorf("job cancelled during process_image: %w", ctx.Err())
	}
}

func handleWriteLog(ctx context.Context, job *queue.Job) error {
	rand.Seed(time.Now().UnixNano())

	log.Printf("write log with payload: %s", job.Payload)

	delay := time.Duration(rand.Intn(10)) * time.Second
	select {
	case <-time.After(delay):
		log.Println("write log successfully after delay")
		return nil

	case <-ctx.Done():
		return fmt.Errorf("job cancelled during write_log: %w", ctx.Err())
	}
}
