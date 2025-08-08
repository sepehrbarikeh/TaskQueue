package postgres

import (
	"TaskQueu/config"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct{
	DB *gorm.DB
}


func New(config config.Db) PostgresDB {
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host,
		config.User,
		config.Pass,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	
	log.Println("✅ Connected to PostgreSQL!")

	return PostgresDB{
		DB: db,
	}
}


func (p *PostgresDB) InsertJobLog(log JobLog) error {
	if log.ID == "" {
		log.ID = uuid.New().String()
	}
	return p.DB.Create(&log).Error
}