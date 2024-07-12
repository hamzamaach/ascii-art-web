package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader sets the HTTP status code and writes the header to the underlying ResponseWriter.
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Log the response status code
func (lrw *loggingResponseWriter) logResponse(logger *log.Logger) {
	logger.Printf("Response %d %s", lrw.statusCode, http.StatusText(lrw.statusCode))
}

// Middleware to log requests
func logMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		defer func() { lrw.ResponseWriter = nil }()

		logger.Printf("%s : %s", r.Method, r.URL)
		next.ServeHTTP(lrw, r)
		lrw.logResponse(logger)
	})
}

// Initializing the Logs file
func LogFile() (*os.File, *log.Logger) {
	// Create a new logger that writes to the file and stdout
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("\033[38;5;196m		Failed to open log file:	\033[0m", err)
		return nil, nil
	}

	// Create a new logger that writes to the file and stdout
	logger := log.New(logFile, "Server : ", log.LstdFlags)

	return logFile, logger
}
