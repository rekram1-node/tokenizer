package tokenizer

import (
	"fmt"
	"strings"

	"github.com/rekram1-node/tokenizer/languages"
)

var (
	// map of  "\t\n\r ,.:?\"!;()" converted to byte for each char
	period           = "."
	defaultSeparator = map[byte]uint8{
		9:  1, // \t
		10: 1, // \n
		13: 1, // \r
		32: 1, // <space>
		44: 1, // ,
		// 46: 1, // .
		58: 1, // L
		63: 1, // ?
		34: 1, // "
		33: 1, // !
		59: 1, // ;
		40: 1, // (
		41: 1, // )
	}
)

// Object used to tokenize
type Tokenizer struct {
	Lanuage         languages.Lanuage
	Separator       map[byte]uint8
	KeepSeparators  bool
	RemoveStopWords bool
}

// Settings for Custom Tokenizer
type Settings struct {
	Lanuage         languages.Lanuage
	KeepSeparators  bool
	RemoveStopWords bool
}

// Default Tokenizer settings
func New() *Tokenizer {
	return &Tokenizer{
		Lanuage:         languages.English,
		Separator:       defaultSeparator,
		KeepSeparators:  false,
		RemoveStopWords: false,
	}
}

func (t *Tokenizer) SetStopWordRemoval(on bool) *Tokenizer {
	t.RemoveStopWords = on
	return t
}

// New Tokenizer Object with your choice of separator, processor, etc
func (s *Settings) Custom(sep string) (*Tokenizer, error) {
	if s.Lanuage.Abbreviations == nil {
		return nil, fmt.Errorf("Invalid Language, missing abbreviations")
	}
	if s.Lanuage.Contractions == nil {
		return nil, fmt.Errorf("Invalid Language, missing contractions")
	}
	if s.Lanuage.StopWords == nil {
		return nil, fmt.Errorf("Invalid Language, missing stop words")
	}
	return &Tokenizer{
		Separator:       convertSeparator(sep),
		KeepSeparators:  s.KeepSeparators,
		RemoveStopWords: s.RemoveStopWords,
		Lanuage:         s.Lanuage,
	}, nil
}

// converts a string into a slice of words
func (t *Tokenizer) TokenizeString(s string) []string {
	/*
		NOTE: this function is messy and full of duplication, however
		as I am priotizing speed and less allocations the messyness is necessary
	*/
	tokens := []string{}
	word := ""
	lastWord := 0
	// isStopWord := false
	for i := 0; i < len(s); i++ {
		if t.Separator[s[i]] != 1 {
			continue
		}

		word = strings.ToLower(strings.TrimSpace(s[lastWord:i]))
		abbr, isAbbreviation := t.Lanuage.Abbreviations[word]

		if strings.Contains(word, period) && isAbbreviation {
			tokens = append(tokens, strings.Split(abbr, " ")...)
			lastWord = i + 1
			if t.KeepSeparators {
				tokens = append(tokens, string(s[i]))
			}
			continue
		} else if strings.Contains(word, period) {
			word = strings.ReplaceAll(word, period, "")
		}

		if !(t.RemoveStopWords && t.Lanuage.StopWords[word] == 1) {
			if word != "" {
				contractions, isContraction := t.Lanuage.Contractions[word]
				if isContraction {
					tokens = append(tokens, strings.Split(contractions, " ")...)
				} else {
					tokens = append(tokens, word)
				}
			}
		}

		lastWord = i + 1

		if t.KeepSeparators {
			tokens = append(tokens, string(s[i]))
		}
	}
	word = strings.ToLower(strings.TrimSpace(s[lastWord:]))
	// if valid
	if !(t.RemoveStopWords && t.Lanuage.StopWords[word] == 1) {
		if word != "" {
			contractions, isContraction := t.Lanuage.Contractions[word]
			if isContraction {
				tokens = append(tokens, strings.Split(contractions, " ")...)
			} else {
				tokens = append(tokens, word)
			}
		}
	}

	return tokens
}

// type TokenizerReader struct {
// 	scanner   *bufio.Scanner
// 	tokenizer *Tokenizer
// 	words     []string
// }

// func NewTokenizerReader(r io.Reader, tokenizer *Tokenizer) *TokenizerReader {
// 	scanner := bufio.NewScanner(r)
// 	return &TokenizerReader{
// 		tokenizer: tokenizer,
// 		scanner:   scanner,
// 	}
// }

// converts file text into words via streamed reading
// perfect for large files
// func (t *TokenizerReader) Scan() bool {
// 	scan := t.scanner.Scan()
// 	word := ""
// 	lastWord := 0
// 	if len(t.scanner.Text()) == 0 {
// 		return scan
// 	}
// 	tokens := make([]string, len(t.scanner.Text())-1)
// 	index := 0
// 	// []string{}
// 	// isStopWord := false
// 	for i := 0; i < len(t.scanner.Text()); i++ {
// 		if t.tokenizer.Separator[t.scanner.Text()[i]] != 1 {
// 			continue
// 		}

// 		word = strings.ToLower(strings.TrimSpace(t.scanner.Text()[lastWord:i]))
// 		abbr, isAbbreviation := t.tokenizer.Lanuage.Abbreviations[word]

// 		if strings.Contains(word, period) && isAbbreviation {
// 			tokens = append(tokens, strings.Split(abbr, " ")...)
// 			lastWord = i + 1
// 			if t.tokenizer.KeepSeparators {
// 				// tokens = append(tokens, string(t.scanner.Text()[i]))
// 				tokens[index] = string(t.scanner.Text()[i])
// 				index++
// 			}
// 			continue
// 		} else if strings.Contains(word, period) {
// 			word = strings.ReplaceAll(word, period, "")
// 		}

// 		if !(t.tokenizer.RemoveStopWords && t.tokenizer.Lanuage.StopWords[word] == 1) {
// 			if word != "" {
// 				contractions, isContraction := t.tokenizer.Lanuage.Contractions[word]
// 				if isContraction {
// 					// tokens = append(tokens, strings.Split(contractions, " ")...)
// 					split := strings.Split(contractions, " ")
// 					for j := 0; j < len(split); j++ {
// 						tokens[index] = split[j]
// 						index++
// 					}
// 				} else {
// 					// tokens = append(tokens, word)
// 					tokens[index] = word
// 					index++
// 				}
// 			}
// 		}

// 		lastWord = i + 1

// 		if t.tokenizer.KeepSeparators {
// 			// tokens = append(tokens, string(t.scanner.Text()[i]))
// 			tokens[index] = string(t.scanner.Text()[i])
// 			index++
// 		}
// 	}
// 	word = strings.ToLower(strings.TrimSpace(t.scanner.Text()[lastWord:]))
// 	// if valid
// 	if !(t.tokenizer.RemoveStopWords && t.tokenizer.Lanuage.StopWords[word] == 1) {
// 		if word != "" {
// 			contractions, isContraction := t.tokenizer.Lanuage.Contractions[word]
// 			if isContraction {
// 				// tokens = append(tokens, strings.Split(contractions, " ")...)
// 				split := strings.Split(contractions, " ")
// 				for j := 0; j < len(split); j++ {
// 					tokens[index] = split[j]
// 					index++
// 				}
// 			} else {
// 				// tokens = append(tokens, word)
// 				tokens[index] = word
// 				index++
// 			}
// 		}
// 	}
// 	t.words = tokens
// 	return scan
// }

// returns tokenize word array
// func (t *TokenizerReader) Words() []string {
// 	return t.words
// }

func convertSeparator(sep string) map[byte]uint8 {
	separators := map[byte]uint8{}
	for i := 0; i < len(sep); i++ {
		separators[sep[i]] = 1
	}

	return separators
}
