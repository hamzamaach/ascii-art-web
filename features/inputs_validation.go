package ascii_art_web

import (
	"net/http"
	"strings"
)

func checkBanner(banner string) bool {
	switch banner {
	case "shadow", "standard", "thinkertoy":
	default:
		return true
	}

	return false
}

// validates if the input contains only printable ASCII characters
func checkValidString(input string) bool {
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, "\n", "")
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			return true
		}
	}
	return false
}

// ValidateInput checks the validity of the input string, banner, and other parameters
func ValidateInput(w http.ResponseWriter, str, banner string) bool {
	if checkBanner(banner) {
		http.Error(w, "404 | Banner not found", http.StatusNotFound)
		return false
	}
	if str == "" {
		http.Error(w, "400 | Bad Request: No input provided", http.StatusBadRequest)
		return false
	}
	if len(str) > 1000 {
		http.Error(w, "413 | The input must contain under 1000 characters.", http.StatusRequestEntityTooLarge)
		return false
	}
	if checkValidString(str) {
		http.Error(w, "400 | Bad Request: Invalid input. The input must contain characters with ASCII values ranging from 32 to 126.", http.StatusBadRequest)
		return false
	}
	return true
}
