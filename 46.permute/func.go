package permutations

import (
	"slices"
)

func permute(nums []int) [][]int {
	stack := [][]int{{}}
	for {
		p := stack[0]
		if len(p) == len(nums) {
			break
		}
		stack = stack[1:]
		for _, n := range nums {
			if !slices.Contains(p, n) {
				pn := append([]int{}, p...)
				stack = append(stack, append(pn, n))
			}
		}
	}
	return stack
}

func permuteBT(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		pm := permuteBT(newNums)
		for _, v := range pm {
			result = append(result, append(v, n))
		}
	}
	return result
}
