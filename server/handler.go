package server

import (
	"TaskQueu/pkg/queue"
	"TaskQueu/repository/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type Handler struct {
	Redis *redis.RedisQueue
}

func NewHandler(redis *redis.RedisQueue) *Handler {
	return &Handler{Redis: redis}
}

type EnqueueRequest struct {
	Payload    string    `json:"payload"`
	QueueName  string    `json:"queue"`
	Type       string    `json:"type"`
	RunAt      time.Time `json:"run_at"`
	MaxRetries int       `json:"max_retries"`
}

func (h *Handler) EnqueueJob(c *fiber.Ctx) error {
	var req EnqueueRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	job := &queue.Job{
		ID:         uuid.New().String(),
		Queue:      req.QueueName,
		Payload:    req.Payload,
		Type:       req.Type,
		RunAt:      req.RunAt,
		MaxRetries: req.MaxRetries,
		RetryCount: 0,
		CreatedAt:  time.Now(),
	}

	if err := h.Redis.Enqueue(c.Context(), req.QueueName, job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to enqueue job",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "job enqueued",
		"id":      job.ID,
	})
}
