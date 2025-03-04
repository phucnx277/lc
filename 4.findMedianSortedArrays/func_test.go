package findmediansortedarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tcs := []struct {
		input    [2][]int
		expected float64
	}{
		{
			input:    [2][]int{{1, 2}, {3, 4}},
			expected: 2.5,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			got := findMedianSortedArrays(tc.input[0], tc.input[1])
			assert.Equal(t, got, tc.expected)
		})
	}
}
