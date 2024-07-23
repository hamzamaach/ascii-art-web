package ascii_art_web

import (
	"fmt"
	"net/http"
	"strings"
)

type Data struct {
	Ascii_art string
	Input     string
	Banner    string
}

// DownloadHandler handles the download of the ASCII art file

// HandleAsciiArt processes the "ascii-art" route
func HandleAsciiArt(w http.ResponseWriter, r *http.Request, tmpl string) {
	str := r.FormValue("string")
	banner := r.FormValue("banner")
	export := r.Form.Has("Download")

	if !ValidateInput(w, str, banner) {
		return
	}

	str = strings.ReplaceAll(str, "\r\n", "\n")

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
	if export {
		filename := "ascii_art" + r.FormValue("filetype")
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Length", fmt.Sprint(len(data.Ascii_art)))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(data.Ascii_art))
	} else {
		RenderTemplate(w, tmpl, data)
	}
}

// MainHandler handles the root route and other casescases
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
