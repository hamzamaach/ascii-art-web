package main

import (
	"fmt"
	"net/http"

	ft "ascii_art_web/features"
)

func main() {
	// Creat a costum file server handler
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", ft.Handler)
	fmt.Println("Starting the server on : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("500 | Internal Server Error :", err)
	}
}
