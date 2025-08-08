package main

import (
	"TaskQueu/config"
	"TaskQueu/pkg/worker"
	"TaskQueu/repository/postgres"
	"TaskQueu/repository/redis"
	"TaskQueu/server"

	"log"
)

func main() {
	cfg, err := config.Config("./")
	if err != nil {
		log.Fatal(err)
	}
	db := postgres.New(cfg.Db)
	db.Migrate()

	redis := redis.NewRedisQueue(cfg.Redis)

	dispatcher := worker.NewDispatcher(&db,redis,"email",5)
	dispatcher.Start()

	userHandler := server.NewHandler(redis)

	server := server.New(cfg.Server,*userHandler)

	server.Serve()

}
