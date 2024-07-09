package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type MyServer struct {
	Router  *http.ServeMux
	Server  *http.Server
	Logger  *log.Logger
	Logfile *os.File
}

func main() {
	fmt.Println("Starting the server...")

	// Create a new server instance
	MyServer := ServerMaker()

	// Close the log file when the program exits
	defer MyServer.Logfile.Close()

	// Start the server
	MyServer.run()

}

func ServerMaker() *MyServer {
	myServer := new(MyServer)

	// Create a the log file and the logger
	logfile, logger := LogFile()

	// Create a new ServeMux instance
	mux := http.NewServeMux()

	// Create a new server instance
	server := &http.Server{
		Addr:    "localhost:1337",
		Handler: logMiddleware(mux, logger), // Use the ServeMux as the server's main handler
	}

	// Assign variables to the server instance
	myServer = &MyServer{
		Router:  mux,
		Server:  server,
		Logger:  logger,
		Logfile: logfile,
	}
	// Map the routes to their respective handlers
	myServer.mapRoutes()

	return myServer
}
