package server

import (
	"TaskQueu/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config      config.Server
	userHandler Handler
	Router *fiber.App
}

func New(config config.Server, userHandler Handler) Server {
	return Server{
		config: config,
		userHandler: userHandler,
		Router: fiber.New(),
	}
}


func (s Server) Serve() {
	app := s.Router

	SetupRoutes(app,s.userHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", s.config.Port)))
}
