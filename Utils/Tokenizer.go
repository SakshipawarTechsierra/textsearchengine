package utils

import (
	"strings"
	"unicode"
)

// Tokenize splits text into words, keeping only letters and numbers
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// analyze performs full preprocessing: tokenize, lowercase, stem
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

// lowercaseFilter converts all tokens to lowercase
func lowercaseFilter(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}
	return tokens
}

// stemmerFilter is a dummy stemmer — you can integrate real stemmer later
func stemmerFilter(tokens []string) []string {
	// Example: just return tokens as-is (no stemming)
	return tokens
}
