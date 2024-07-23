package ascii_art_web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ProcessInput processes the input string, reads the banner, and produces the ASCII art
func ProcessInput(w http.ResponseWriter, input, banner string) (string, error) {
	inputWithoutNewLines := strings.ReplaceAll(input, "\n", "")
	if len(inputWithoutNewLines) == 0 {
		return input, nil
	}

	charactersMap, err := readBanner(banner)
	if err != nil {
		return "", err
	}

	return drawASCIIArt(charactersMap, input), nil
}

// readBanner reads the banner file and returns the ASCII art characters in map [rune]string and err if it occurs
func readBanner(banner string) (map[rune][]string, error) {
	data, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Printf("500 | Internal Server Error: Unable to read banner file: %s\n", err)
		return nil, err
	}
	stringData := string(data[1:])
	if banner == "thinkertoy" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}
	content := strings.Split(stringData, "\n\n")
	charactersMap := convertTocharacterMap(content)
	return charactersMap, nil
}

// DrawASCIIArt draws ASCII art and colorizes specific substrings
// Render the ASCII art based on the character map and the input lines
func drawASCIIArt(characterMatrix map[rune][]string, input string) string {
	result := ""
	splittedInput := strings.Split(input, "\n")

	for _, val := range splittedInput {
		if val == "" {
			result += "\n"
		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					result += characterMatrix[k][j]
				}
				result += "\n"
			}
		}
	}
	return result
}

// Convert content array to a character mapping ASCII characters to their line representations
func convertTocharacterMap(content []string) map[rune][]string {
	charactersMap := map[rune][]string{}
	for i, val := range content {
		charactersMap[rune(32+i)] = strings.Split(val, "\n")
	}
	return charactersMap
}
