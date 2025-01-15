package generateparenthesis

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tcs := []struct {
		input    int
		expected []string
	}{
		{
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			got := generateParenthesis(tc.input)
			slices.Sort(got)
			slices.Sort(tc.expected)
			assert.ElementsMatch(t, got, tc.expected)
		})
	}
}
