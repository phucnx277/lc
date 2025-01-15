package generateparenthesis

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	result := []string{"("}
	for {
		first := result[0]
		if len(first) == n*2 {
			break
		}
		result = result[1:]
		open, closed := getCounts(first)
		if n-open > 0 {
			result = append(result, first+"(")
		}
		if n-closed > 0 && open > closed {
			result = append(result, first+")")
		}
	}
	return result
}

func getCounts(pattern string) (int, int) {
	open, closed := 0, 0
	for _, r := range pattern {
		if r == '(' {
			open += 1
		} else if r == ')' {
			closed += 1
		}
	}
	return open, closed
}
