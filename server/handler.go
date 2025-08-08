package server

import (
	"TaskQueue/pkg/queue"
	"TaskQueue/repository/redis"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	// Validate required fields
	if req.QueueName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "queue name is required",
		})
	}

	if req.Payload == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "payload is required",
		})
	}

	if req.Type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "job type is required",
		})
	}

	// Set default values
	if req.MaxRetries == 0 {
		req.MaxRetries = 3
	}

	if req.RunAt.IsZero() {
		req.RunAt = time.Now()
	}

	// Set default type if not provided
	if req.Type == "" {
		req.Type = req.QueueName
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
		"message": "job enqueued successfully",
		"id":      job.ID,
		"queue":   req.QueueName,
	})
}
