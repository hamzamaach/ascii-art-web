package main

import (
	"fmt"
	"net/http"

	ft "ascii_art_web/features"
)

func main() {
	http.HandleFunc("/", ft.Handler)
	fmt.Println("Starting the server on : http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
