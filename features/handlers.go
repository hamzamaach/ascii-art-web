package ascii_art_web

import (
	"net/http"
	"os"
)

type Data struct {
	Ascii_art string
	Input     string
	Banner    string
}

// DownloadHandler handles the download of the ASCII art file
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" {
		http.Error(w, "400 | Bad Request: No file specified", http.StatusBadRequest)
		return
	}
	content, err := os.ReadFile(file)
	if err != nil {
		http.Error(w, "404 | Not Found: File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+file)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(content)
}

// HandleAsciiArt processes the "ascii-art" route
func HandleAsciiArt(w http.ResponseWriter, r *http.Request, tmpl string) {
	str := r.FormValue("string")
	banner := r.FormValue("banner")
	export := r.FormValue("export")

	if !ValidateInput(w, str, banner) {
		return
	}

	asciiArt, err := ProcessInput(w, str, banner)
	if err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		return
	}

	data := Data{
		Ascii_art: asciiArt,
		Input:     str,
		Banner:    banner,
	}

	switch export {
	case "download":
		filename := "ascii_art.txt"
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(data.Ascii_art))
	default:
		RenderTemplate(w, tmpl, data)
	}
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

		RenderTemplate(w, tmpl, Data{})

	case "/ascii-art":
		if r.Method != http.MethodPost {
			http.Error(w, "405 | Method Not Allowed: Use POST", http.StatusMethodNotAllowed)
			return
		}
		HandleAsciiArt(w, r, tmpl)

	case "/about":
		RenderTemplate(w, "about.html", Data{})

	default:
		http.NotFound(w, r)
	}
}
