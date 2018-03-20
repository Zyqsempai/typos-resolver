package typos

import (
	"bytes"
)

func MakeReverse(s string, toLang string) (reverse string) {
	switch toLang {
	case "en":
		reverse = substituteRune(s, ruToEn)
	case "ru":
		reverse = substituteRune(s, enToRu)
	default: // fallback to "en" if lang not found
		reverse = substituteRune(s, ruToEn)
	}

	return reverse
}

func substituteRune(s string, sub map[rune]rune) string {
	var buf bytes.Buffer
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf.WriteRune(d)
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func SubstituteTypos(s string, toLang string) (sub string) {
	switch toLang {
	case "en":
		sub = substituteRune(s, ruEnIdenticChars)
	case "ru":
		sub = substituteRune(s, enRuIdenticChars)
	default: // fallback to "en" if lang not found
		sub = substituteRune(s, ruEnIdenticChars)
	}

	return sub
}
