package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Middleware to log requests
func logMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Create a new file for logs
func LogFile() (*os.File, *log.Logger) {
	// Create a new logger that writes to the file and stdout
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return nil, nil
	}
	// Create a new logger that writes to the file and stdout
	logger := log.New(io.MultiWriter(logFile, os.Stdout), "Server: ", log.LstdFlags)

	return logFile, logger
}
