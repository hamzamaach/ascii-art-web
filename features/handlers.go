package ascii_art_web

import (
	"net/http"
)

// HandleAsciiArt processes the "ascii-art" route
func HandleAsciiArt(w http.ResponseWriter, r *http.Request, tmpl string) {
	str := r.FormValue("string")
	banner := r.FormValue("banner")
	if CheckValidInput(str) || CheckBanner(banner) {
		http.Error(w, "400 | Bad Request: Invalid input or banner type", http.StatusBadRequest)
		return
	}
	data := ProcessInput(str, banner)
	RenderTemplate(w, tmpl, data)
}

// MainHandler handles the root route and other cases
func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := "index.html"

	switch r.URL.Path {
	case "/":
		RenderTemplate(w, tmpl, nil)
	case "/ascii-art":
		HandleAsciiArt(w, r, tmpl)
	default:
		http.NotFound(w, r)
	}
}
