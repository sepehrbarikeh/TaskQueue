package main

import (
	"TaskQueue/config"
	"TaskQueue/pkg/worker"
	"TaskQueue/repository/postgres"
	"TaskQueue/repository/redis"
	"TaskQueue/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Config("./")
	if err != nil {
		log.Fatal(err)
	}
	db := postgres.New(cfg.Db)
	db.Migrate()

	redis := redis.NewRedisQueue(cfg.Redis)
	if redis == nil {
		log.Fatal("Failed to connect to Redis")
	}

	dispatcher := worker.NewDispatcher(&db, redis, "send_email", 5)
	dispatcher.Start()

	userHandler := server.NewHandler(redis)

	server := server.New(cfg.Server, *userHandler)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		server.Serve()
	}()

	<-quit
	log.Println("Shutting down...")
	dispatcher.Stop()
	log.Println("Server stopped")
}
