package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/quenyu/pr-reviewer/internal/config"
	"github.com/quenyu/pr-reviewer/internal/server"
)

func main() {
	cfg := config.Load()

	srv := server.New(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	<-ctx.Done()
	log.Println("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
