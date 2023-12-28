package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSums(t *testing.T) {
	tests := []struct {
		expected   []int
		inputArray []int
		target     int
	}{
		{
			expected:   []int{0, 1},
			inputArray: []int{2, 7, 11, 15},
			target:     9,
		},
		{
			expected:   []int{1, 2},
			inputArray: []int{3, 2, 4},
			target:     6,
		},
		{
			expected:   []int{0, 1},
			inputArray: []int{3, 3},
			target:     6,
		},
	}

	for _, test := range tests {
		actual := twoSum(test.inputArray, test.target)
		assert.Equal(t, test.expected, actual, "Slices are not equal")
	}
}
