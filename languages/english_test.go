package languages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected bool
	}{
		{
			name:     "English Stop Word",
			word:     "above",
			expected: true,
		},
		{
			name:     "Not English Stop Word",
			word:     "church",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, isStopWord := English.StopWords[test.word]
			assert.Equal(t, test.expected, isStopWord)
		})
	}
}
