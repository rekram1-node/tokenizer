package tokenizer

import (
	"strings"
)

var (
	// map of  "\t\n\r ,.:?\"!;()" converted to byte for each char
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

// Object used to tokenize
type Tokenizer struct {
	Separator      map[byte]uint8
	KeepSeparators bool

	// defaults to removing stopwords
	WordChecker WordProcessor
}

// function to determine if a word is a valid token
type WordProcessor func(word string) bool

// Default Tokenizer settings
func New() *Tokenizer {
	return &Tokenizer{
		Separator: defaultSeparator,
		WordChecker: func(word string) bool {
			return true
		},
	}
}

// New Tokenizer Object with your choice of separator, processor, etc
func Custom(sep string, keepSep bool, wordProcessor WordProcessor) *Tokenizer {
	return &Tokenizer{
		Separator:      convertSeparator(sep),
		KeepSeparators: keepSep,
		WordChecker:    wordProcessor,
	}
}

// converts a string into a slice of words
func (t *Tokenizer) TokenizeByWord(s string) []string {
	s = strings.TrimSpace(s)
	tokens := []string{}
	lastWord := 0
	for i := range s {
		if t.Separator[s[i]] != 1 {
			continue
		}

		word := s[lastWord:i]
		if t.validWord(word) {
			tokens = append(tokens, s[lastWord:i])
		}
		lastWord = i + 1

		if t.Separator[s[i]] == 1 && t.KeepSeparators {
			word := strings.TrimSpace(string(s[i]))
			if t.validWord(word) {
				tokens = append(tokens, word)
			}
		}
	}
	word := strings.TrimSpace(s[lastWord:])
	if t.validWord(word) {
		tokens = append(tokens, s[lastWord:])
	}

	return tokens
}

// convert a string into slices of word slices
// Coming soon
func (t *Tokenizer) TokenizeBySentence() {

}

func convertSeparator(sep string) map[byte]uint8 {
	separators := map[byte]uint8{}
	for i := 0; i < len(sep); i++ {
		separators[sep[i]] = 1
	}

	return separators
}

func (t *Tokenizer) validWord(word string) bool {
	return word != "" && t.WordChecker(word)
}
