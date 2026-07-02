package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmvsnx/inventory-manegement/internal/config"
	"github.com/dmvsnx/inventory-manegement/internal/database"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app fiber.App
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		app: *fiber.New(),
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	_, err := database.NewDB(s.cfg)
	if err != nil {
		return fmt.Errorf("init database: %w", err)
	}

	addr := fmt.Sprintf(":%s", s.cfg.AppPort)
	go func() {
		log.Printf("Server running on %s", addr)
		if err := s.app.Listen(addr); err != nil {
			log.Printf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("server shutdown: %w", err)
	}

	log.Println("Server shutdown successfully")
	return nil
}
