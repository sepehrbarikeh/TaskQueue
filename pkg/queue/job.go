package queue


import "time"

type Job struct {
    ID          string    `json:"id"`
    Queue       string    `json:"queue"`
    Payload     string    `json:"payload"` // به صورت JSON string
    RunAt       time.Time `json:"run_at"`
    MaxRetries  int       `json:"max_retries"`
    RetryCount  int       `json:"retry_count"`
    CreatedAt   time.Time `json:"created_at"`
}
