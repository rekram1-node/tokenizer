package tokenizer

import (
	"strings"
)

const (
	defaultRegex = "\t\n\r ,.:?\"!;()"
)

var (
	defaultSeparator = map[byte]uint8{
		9:  1,
		10: 1,
		13: 1,
		32: 1,
		44: 1,
		46: 1,
		58: 1,
		63: 1,
		34: 1,
		33: 1,
		59: 1,
		40: 1,
		41: 1,
	}
)

type Tokenizer struct {
	Separator      map[byte]uint8
	KeepSeparators bool
}

func New() *Tokenizer {
	return &Tokenizer{
		Separator: defaultSeparator,
	}
}

func Custom(sep string) *Tokenizer {
	return &Tokenizer{
		Separator: convertSeparator(sep),
	}
}

func (t *Tokenizer) TokenizeByWord(s string) []string {
	s = strings.TrimSpace(s)
	tokens := []string{}
	lastWord := 0
	for i := 0; i < len(s); i++ {
		if t.Separator[s[i]] != 1 {
			continue
		}

		word := s[lastWord:i]
		if word != "" {
			tokens = append(tokens, s[lastWord:i])
		}
		lastWord = i + 1

		if t.Separator[s[i]] == 1 && t.KeepSeparators {
			word := strings.TrimSpace(string(s[i]))
			if word != "" {
				tokens = append(tokens, word)
			}
		}
	}
	word := strings.TrimSpace(s[lastWord:])
	if word != "" {
		tokens = append(tokens, s[lastWord:])
	}

	return tokens
}

func (t *Tokenizer) TokenizeBySentence() {}

func convertSeparator(sep string) map[byte]uint8 {
	separators := map[byte]uint8{}
	for i := 0; i < len(sep); i++ {
		separators[sep[i]] = 1
	}

	return separators
}
