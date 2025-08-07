package main

import (
	"TaskQueu/config"
	"TaskQueu/repository/postgres"
	"TaskQueu/repository/redis"
	"fmt"

	"log"
)

func main() {
	cfg,err:= config.Config("./")
	if err != nil {
		log.Fatal(err)
	}
	db := postgres.New(cfg.Db)
	db.Migrate()


	redis := redis.NewRedisQueue(cfg.Redis)

	fmt.Println(redis)

}