package utils

import (
	"regexp"
	"strings"
)

func Tokenize(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	return re.FindAllString(strings.ToLower(text), -1)
}
