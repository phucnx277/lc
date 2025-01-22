package combinationsum

import (
	"slices"
	"strconv"
)

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	resultMap := make(map[string]struct{})
	stack := [][]int{}
	for _, v := range candidates {
		if v <= target {
			stack = append(stack, []int{v})
		}
	}
	for len(stack) > 0 {
		coms := stack[0]
		stack = stack[1:]
		sum := calSum(coms)
		if sum == target {
			slices.Sort(coms)
			key := getKey(coms)
			_, exist := resultMap[key]
			if !exist {
				resultMap[key] = struct{}{}
				result = append(result, coms)
			}
			continue
		}
		for _, v := range candidates {
			if sum+v <= target {
				stack = append(stack, append([]int{v}, coms...))
			}
		}
	}
	return result
}

func calSum(coms []int) int {
	sum := 0
	for _, v := range coms {
		sum += v
	}
	return sum
}

func getKey(coms []int) string {
	key := ""
	for _, v := range coms {
		key += strconv.Itoa(v)
	}
	return key
}
