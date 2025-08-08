package postgres

import (
	"time"
)

// مدل جدول job_logs
type JobLog struct {
	ID         string `gorm:"primaryKey;type:uuid"`
	JobID      string `gorm:"type:uuid;not null"`
	Queue      string `gorm:"not null"`
	Status     string `gorm:"not null"`
	Payload    string `gorm:"type:text"`
	RetryCount int
	Error      string    `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
