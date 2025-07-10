package utils

func MatchesAny(tokens, keywords []string) bool {
	tokenMap := make(map[string]bool)
	for _, token := range tokens {
		tokenMap[token] = true
	}
	for _, keyword := range keywords {
		if tokenMap[keyword] {
			return true
		}
	}
	return false
}
