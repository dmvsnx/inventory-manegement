package main

import (
	"log"

	"github.com/dmvsnx/inventory-manegement/internal/app"
	"github.com/dmvsnx/inventory-manegement/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	server := app.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}