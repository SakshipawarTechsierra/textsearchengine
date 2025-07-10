package utils

var stopwords = map[string]bool{
	"the": true, "is": true, "in": true, "of": true,
	"and": true, "a": true, "to": true, "on": true,
}

func Filter(tokens []string) []string {
	var filtered []string
	for _, token := range tokens {
		if !stopwords[token] {
			filtered = append(filtered, token)
		}
	}
	return filtered
}
