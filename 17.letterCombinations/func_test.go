package lettercombinations

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tcs := []struct {
		input    string
		expected []string
	}{
		{
			input:    "672",
			// expected: []string{"mpa", "mqa", "mra", "msa", "npa", "nqa", "nra", "nsa", "opa", "oqa", "ora", "osa", "mpb", "mqb", "mrb", "msb", "npb", "nqb", "nrb", "nsb", "opb", "oqb", "orb", "osb", "mpc", "mqc", "mrc", "msc", "npc", "nqc", "nrc", "nsc", "opc", "oqc", "orc", "osc"},
			expected: letterCombinationsBF("672"),
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			got := letterCombinations(tc.input)
			slices.Sort(got)
			slices.Sort(tc.expected)
			assert.ElementsMatch(t, got, tc.expected)
		})
	}
}
