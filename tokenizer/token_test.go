package tokenizer

import (
	"testing"

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
				"The", "world", "is", "a", "wonderful", "place", "There", "are", "many", "places",
				"in", "this", "world", "This", "place", "is", "wonderful",
			},
		},
		{
			name: "Tokenize: Keep Separators",
			t:    Custom(defaultRegex, true, func(word string) bool { return true }),
			input: `
			The world is a wonderful place...
			There are many places in this world!!!	
			
			This place is wonderful.
				`,
			expectedSlice: []string{
				"The", "world", "is", "a", "wonderful", "place", ".", ".", ".", "There", "are", "many", "places",
				"in", "this", "world", "!", "!", "!",
				"This", "place", "is", "wonderful", ".",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			slice := test.t.TokenizeByWord(test.input)
			assert.Equal(t, test.expectedSlice, slice)
		})
	}
}
