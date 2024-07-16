package ascii_art_web

import (
	"net/http"
	"strings"
)

// Convert content array to a character mapping ASCII characters to their line representations
func ConvertTocharacterMap(content []string) map[rune][]string {
	charactersMap := map[rune][]string{}
	for i, val := range content {
		charactersMap[rune(32+i)] = strings.Split(val, "\n")
	}
	return charactersMap
}

// ProcessInput processes the input string, reads the banner, and produces the ASCII art
func ProcessInput(w http.ResponseWriter, input, banner string) (string, error) {
	splittedInput := strings.Split(input, "\r\n")

	charactersMap, err := ReadBanner(banner, w)
	if err != nil {
		return "", err
	}

	result := DrawASCIIArt(charactersMap, splittedInput)
	return strings.Join(result, "\n"), nil
}
