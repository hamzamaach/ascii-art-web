package ascii_art_web

import (
	"fmt"
	"os"
	"strings"
)

func ReadBanner(banner string) map[rune][]string {
	data, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	stringData := string(data[1:])
	if banner == "thinkertoy" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}
	content := strings.Split(stringData, "\n\n")
	charactersMap := ConvertTocharacterMap(content)
	return charactersMap
}
