package ascii_art_web

import "strings"

func CheckBanner(banner string) bool {
	switch banner {
	case "shadow", "standard", "thinkertoy":
	default:
		return true
	}

	return false
}

// validates if the input contains only printable ASCII characters
func CheckValidInput(input string) bool {
	input = strings.ReplaceAll(input, "\r", "")
	input = strings.ReplaceAll(input, "\n", "")
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			return true
		}
	}
	return false
}
