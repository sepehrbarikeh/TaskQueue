package server

import (
	"TaskQueue/config"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config      config.Server
	userHandler Handler
	Router      *fiber.App
}

func New(config config.Server, userHandler Handler) Server {
	return Server{
		config:      config,
		userHandler: userHandler,
		Router:      fiber.New(),
	}
}

func (s Server) Serve() {
	app := s.Router

	SetupRoutes(app, s.userHandler)

	// Graceful shutdown
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", s.config.Port)); err != nil {
			log.Fatal("Server error:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited")
}
