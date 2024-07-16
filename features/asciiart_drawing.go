package ascii_art_web

import (
	"strings"
)

// Check if there are any non-empty lines in the input lines array
func checkEmptyLines(splittedInput []string) bool {
	for _, line := range splittedInput {
		if line != "" {
			return false
		}
	}
	return true
}

// DrawASCIIArt draws ASCII art and colorizes specific substrings
func DrawASCIIArt(
	charactersMap map[rune][]string,
	splittedInput []string,
) []string {
	var result []string

	// check if the input contain only new lines
	emptyLines := checkEmptyLines(splittedInput)
	if emptyLines {
		newLines := strings.Repeat("\n", len(splittedInput)-1)
		result = append(result, newLines)
		return result
	}

	for _, inputLine := range splittedInput {
		var resultLine strings.Builder
		if inputLine == "" {
			result = append(result, "\n")
			continue
		}

		// Draw each character of the line in ASCII art format
		for line := 0; line < 8; line++ {
			for _, char := range inputLine {
				resultLine.WriteString(charactersMap[char][line])
			}
			resultLine.WriteString("\n")
		}
		result = append(result, resultLine.String())
	}
	return result
}
