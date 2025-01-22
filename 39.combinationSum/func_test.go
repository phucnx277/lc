package combinationsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tcs := []struct {
		inputCandidates []int
		inputTarget     int
		expected        [][]int
	}{
		{
			inputCandidates: []int{7,3,2},
			inputTarget:     18,
			expected:        [][]int{{2,2,7,7},{2,3,3,3,7},{2,2,2,2,3,7},{3,3,3,3,3,3},{2,2,2,3,3,3,3},{2,2,2,2,2,2,3,3},{2,2,2,2,2,2,2,2,2}},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			got := combinationSum(tc.inputCandidates, tc.inputTarget)
			assert.ElementsMatch(t, got, tc.expected)
		})
	}
}
