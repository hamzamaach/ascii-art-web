package ascii_art_web

import (
	"net/http"
)

// HandleAsciiArt processes the "ascii-art" route
func HandleAsciiArt(w http.ResponseWriter, r *http.Request, tmpl string) {
	str := r.FormValue("string")
	banner := r.FormValue("banner")

	if CheckBanner(banner) {
		http.Error(w, "404 | Bad Request: Banner not found", http.StatusNotFound)
		return
	}
	if str == "" {
		http.Error(w, "400 | Bad Request: No input provided", http.StatusBadRequest)
		return
	}

	if CheckValidInput(str) {
		http.Error(w, "400 | Bad Request: Invalid input. The input must contain characters with ASCII values ranging from 32 to 126.", http.StatusBadRequest)
		return
	}

	data, err := ProcessInput(w, str, banner)
	if err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		return
	}
	RenderTemplate(w, tmpl, data)
}

// MainHandler handles the root route and other cases
func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := "index.html"

	switch r.URL.Path {
	case "/":
		if r.Method != http.MethodGet {
			http.Error(w, "405 | Method Not Allowed: Use GET", http.StatusMethodNotAllowed)
			return
		}

		RenderTemplate(w, tmpl, nil)
	case "/ascii-art":
		if r.Method != http.MethodPost {
			http.Error(w, "405 | Method Not Allowed: Use POST", http.StatusMethodNotAllowed)
			return
		}
		HandleAsciiArt(w, r, tmpl)
	default:
		http.NotFound(w, r)
	}
}
