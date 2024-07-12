package main

import (
	"html/template"
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

	return myServer
}

// Map the routes to their respective handlers
func (srv *MyServer) mapRoutes() {
	// Get the list of handlers
	handlers := getHandlers()

	// Register the handlers with the ServeMux
	for pattern, handler := range handlers {
		srv.Router.HandleFunc(pattern, handler)
	}
	srv.Router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func getHandlers() map[string]http.HandlerFunc {
	// Create a map to store the handlers
	handlers := make(map[string]http.HandlerFunc)

	// Define the handler functions
	lpHandler := func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed!\n", http.StatusMethodNotAllowed)
			return
		}

		// Check if the request path exists "/*"
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/404", http.StatusTemporaryRedirect)
			return
		}

		// Parse the template file
		tmpl, err := template.ParseFiles("../templates/index.html")
		if err != nil {
			http.Error(w, "Template not found!\n", http.StatusBadGateway)
			return
		}
		// Execute the template with the data
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error executing the template!\n", http.StatusInternalServerError)
			return
		}
	}

	asciiArtHandler := func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed!\n", http.StatusMethodNotAllowed)
			return
		}
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data!\n", http.StatusBadRequest)
			return
		}

		// Get the form values
		// str := r.FormValue("String")
		// banner := r.FormValue("banner")

		// ******************************************************
		// if CheckValidInput(str) || CheckBanner(banner) {
		// 	http.Error(w, "400 | Bad Request: Invalid input or banner type", http.StatusBadRequest)
		// 	return
		// }
		// data := ProcessInput(str, banner)
		// ******************************************************
	}

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("\n\n\n										    ❤️ Welcome to our first Gofer powered Server ❤️ 				\n\n" +
			"												     Credit :\n\n" +
			"											 Ismail Bentour | Hamza Maach\n\n" +
			"												@2024 Zone01 Oujda \n"))
	}

	notFoundHandler := func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}

	// Add the handlers to the map
	handlers["/"] = lpHandler
	handlers["/about"] = aboutHandler
	handlers["/ascii-art"] = asciiArtHandler
	handlers["/404"] = notFoundHandler

	return handlers
}
