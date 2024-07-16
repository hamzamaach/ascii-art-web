package ascii_art_web

import (
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
func ProcessInput(input, banner string) string {
	splittedInput := strings.Split(input, "\r\n")

	charactersMap := ReadBanner(banner)

	result := DrawASCIIArt(charactersMap, splittedInput)
	return strings.Join(result, "\n")
}
