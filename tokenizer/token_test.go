package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			t: &Tokenizer{
				Separator:      convertSeparator(defaultRegex),
				KeepSeparators: true,
			},
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
