package main

import (
	"context"
	"html/template"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// Map the routes to their respective handlers
func (srv *MyServer) mapRoutes() {
	// Define the handler functions
	lpHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template not found!", http.StatusNotFound)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error executing template!", http.StatusInternalServerError)
			return
		}
	}

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to my first Gofer powered Server\n"))
	}

	// Register the handlers with the ServeMux
	srv.Router.HandleFunc("/", lpHandler)
	srv.Router.HandleFunc("/about", aboutHandler)
}

// Start the server
func (srv *MyServer) run() {
	srv.Logger.Println("Starting the server on :", srv.Server.Addr)

	// Create a context to handle shutdown signals
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Start the server in a separate goroutine
	go func() {
		if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.Logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for a shutdown signal
	<-ctx.Done()

	// Initiate a graceful shutdown
	srv.Logger.Println("Shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Server.Shutdown(shutdownCtx); err != nil {
		srv.Logger.Fatalf("Failed to shut down server gracefully: %v", err)
	}

	srv.Logger.Println("Server stopped gracefully")
}
