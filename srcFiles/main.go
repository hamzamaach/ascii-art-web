package main

import (
	"fmt"
)

func main() {
	fmt.Println("\033[32m		Starting the server...\033[0m")

	// Create a new server instance
	MyServer := ServerMaker()

	// Map the routes to their respective handlers
	MyServer.mapRoutes()

	// Close the log file when the program exits
	defer MyServer.Logfile.Close()

	// Start the server
	MyServer.run()
}
