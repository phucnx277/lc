package lettercombinations

import "fmt"

// 2-loop
func letterCombinationsBF(digits string) []string {
	dl := len(digits)
	if dl == 0 {
		return []string{}
	}
	lettersMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	out := lettersMap[string(digits[0])]
	for i := 1; i < dl; i++ {
		letters := lettersMap[string(digits[i])]
		m, n := len(out), len(letters)
		output := make([]string, m*n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				output[i*n+j] = out[i] + letters[j]
			}
		}
		out = output
	}
	return out
}

// back tracking
func letterCombinations(digits string) []string {
	dl := len(digits)
	if dl == 0 {
		return []string{}
	}
	lettersMap := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	out := []string{}
	backtrack := func(i int, c string) {}
	backtrack = func(index int, curStr string) {
		if len(curStr) == dl {
			out = append(out, curStr)
			return
		}
		chars := lettersMap[string(digits[index])]
		fmt.Println(string(digits[index]))
		for _, c := range chars {
			backtrack(index+1, curStr+c)
		}
	}
	backtrack(0, "")
	return out
}
