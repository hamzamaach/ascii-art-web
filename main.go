package main

import (
	"net/http"

	ft "ascii_art_web/features"
)

func main() {
	http.HandleFunc("/", ft.Handler)
	http.ListenAndServe("localhost:8080", nil)
}
