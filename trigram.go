package tinygram

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"golang.org/x/text/runes"
)

func generateTrigrams(s string) []string {
	s = normalizeText(s)

	var trigrams []string
	if len(s) < 3 {
		return []string{s} // handle very short strings
	}
	for i := 0; i <= len(s)-3; i++ {
		trigram := s[i : i+3]
		trigrams = append(trigrams, trigram)
	}
	return trigrams
}

func normalizeText(input string) string {
	s := strings.ToLower(strings.TrimSpace(input))

	// Unicode normalize and remove diacritics
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)
	s, _, _ = transform.String(t, s)

	// Remove punctuation
	rePunct := regexp.MustCompile(`[!"#$%'\(\)\*\+,\-./:;<=>?@\[\\\]^_` + "`" + `{|}~]`)
	s = rePunct.ReplaceAllString(s, " ")

	// Collapse multiple spaces
	reSpace := regexp.MustCompile(`\s+`)
	s = reSpace.ReplaceAllString(s, " ")
	return s
}
