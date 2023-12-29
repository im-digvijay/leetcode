package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSums(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{
			input:    121,
			expected: true,
		},
		{
			input:    -121,
			expected: false,
		},
		{
			input:    10,
			expected: false,
		},
	}

	for _, test := range tests {
		actual := isPalindrome(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
