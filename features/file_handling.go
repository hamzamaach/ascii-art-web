package ascii_art_web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ReadBanner(banner string, w http.ResponseWriter) (map[rune][]string, error) {
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
	charactersMap := ConvertTocharacterMap(content)
	return charactersMap, nil
}
