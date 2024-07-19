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

// DrawASCIIArt draws ASCII art and colorizes specific substrings
// Render the ASCII art based on the character map and the input lines
func drawASCIIArt(characterMatrix map[rune][]string, input string) string {
	result := ""
	splittedInput := strings.Split(input, "\r\n")

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

// ProcessInput processes the input string, reads the banner, and produces the ASCII art
func ProcessInput(w http.ResponseWriter, input, banner string) (string, error) {
	inputWithoutNewLines := strings.ReplaceAll(input, "\r\n", "")
	if len(inputWithoutNewLines) == 0 {
		return strings.Repeat("\n", strings.Count(input, "\r\n")), nil
	}

	charactersMap, err := ReadBanner(banner, w)
	if err != nil {
		return "", err
	}

	return drawASCIIArt(charactersMap, input), nil
}
