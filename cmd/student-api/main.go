package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shantanuj2002/students-api/internal/config"
	"github.com/shantanuj2002/students-api/internal/http/handler/student"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Set up router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New())
	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to student api"))
	// })

	// Create server
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// Channel to listen for interrupt/terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run server in a goroutine
	go func() {
		fmt.Printf("Server started at %s\n", cfg.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Block until signal is received
	<-stop
	slog.Info("\nShutting down server...")

	// Graceful shutdown with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown Failed: %v", err)
	}

}
