package permutations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tcs := []struct {
		inputNums []int
		expected  [][]int
	}{
		{
			inputNums: []int{1, 2, 3},
			expected:  [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			got1 := permuteBT(tc.inputNums)
			got2 := permute(tc.inputNums)
			assert.ElementsMatch(t, got1, tc.expected)
			assert.ElementsMatch(t, got2, tc.expected)
		})
	}
}
