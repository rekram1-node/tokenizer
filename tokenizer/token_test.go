package tokenizer

import (
	"testing"

	"github.com/rekram1-node/tokenizer/languages"
	"github.com/stretchr/testify/assert"
)

const (
	defaultRegex = "\t\n\r ,.:?\"!;()"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name          string
		t             *Tokenizer
		input         string
		expectedSlice []string
	}{
		{
			name: "Tokenize: Expect correct outcome",
			t:    New(),
			input: `
			The world is a wonderful place...
			There are many places in this world!!!	

			
			This place is wonderful.
				`,
			expectedSlice: []string{
				"the", "world", "is", "a", "wonderful", "place", "there", "are", "many", "places",
				"in", "this", "world", "this", "place", "is", "wonderful",
			},
		},
		{
			name: "Tokenize: Replace abbreviations",
			t:    New(),
			input: `
			The world is a wonderful place...
			There are many places in this world!!!	

			The month is feb.

			The hr. is now, 
			
			This place is wonderful.
				`,
			expectedSlice: []string{"the", "world", "is", "a", "wonderful", "place", "there", "are", "many", "places", "in", "this", "world", "the", "month", "is", "february", "the", "hour", "is", "now", "this", "place", "is", "wonderful"},
		},
		{
			name: "Tokenize: Replace Contraction",
			t:    New(),
			input: `
			The world is a wonderful place...
			There are many places in this world!!!	

			I can't believe they've done this

			I'd love to haven't
			
			This place is wonderful.
				`,
			expectedSlice: []string{"the", "world", "is", "a", "wonderful", "place", "there", "are", "many", "places", "in", "this", "world", "i", "cannot", "believe", "they", "have", "done", "this", "i", "would", "love", "to", "have", "not", "this", "place", "is", "wonderful"},
		},
		{
			name: "Tokenize: Keep Separators",
			t: &Tokenizer{
				Separator:       convertSeparator(defaultRegex),
				KeepSeparators:  true,
				RemoveStopWords: true,
				Lanuage:         languages.English,
			},
			input: `
The world is a wonderful place...
There are many places in this world!!!	

This place is wonderful.
				`,
			expectedSlice: []string{"\n", " ", "world", " ", " ", " ", "wonderful", " ", "place", ".", ".", ".", "\n", " ", " ", " ", "places", " ", " ", " ", "world", "!", "!", "!", "\t", "\n", "\n", " ", "place", " ", " ", "wonderful", ".", "\n", "\t", "\t", "\t", "\t"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			slice := test.t.TokenizeString(test.input)
			assert.Equal(t, test.expectedSlice, slice)
		})
	}
}
