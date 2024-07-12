package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// Start the server
func (srv *MyServer) run() {
	fmt.Println("\033[38;5;208m\n	Welcome to our first Gofer powered Server\033[0m\n\033[32m	 Starting the server on :",
		"\033[0m \033[38;5;27m", srv.Server.Addr, "	\033[0m")
	srv.Logger.Println("Starting the server on :", srv.Server.Addr)

	// Create a context to handle shutdown signals
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start the server in a separate goroutine
	go func() {
		if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("\033[38;5;196m\n	Failed to start server : %v	\033[0m\n", err)
			srv.Logger.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	// Wait for a shutdown signal
	<-ctx.Done()

	// Initiate a graceful shutdown
	fmt.Println("\033[38;5;202m\n		Shutdown signal received !\n		Shutting down the server...	\033[0m")
	srv.Logger.Println("Shutdown signal received! Shutting down the server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("\033[38;5;196m	Failed to shut down server gracefully: %v	\033[0m\n", err)
		srv.Logger.Fatalf("Failed to shut down server gracefully: %v\n", err)
	}

	fmt.Println("\033[32m\n		Server stopped gracefully ! 😎\033[0m")
	srv.Logger.Println("Server stopped gracefully")
}
