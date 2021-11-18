package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
	//nums := []int{1,2,3,4,5}
	//newNums := append(nums[:3], nums[4:]...)
	//fmt.Println("nums:", nums, "newNums:", newNums)
}

func permute(nums []int) [][]int {
	results := [][]int{}
	var backtrack func([]int, []int)

	backtrack = func(nums []int, result []int) {
		if len(result) == len(nums) {
			newResult := append([]int{}, result...)
			results = append(results, newResult)
			return
		}

		for _, v := range nums {
			if contains(result, v) {
				continue
			}
			result = append(result, v)
			backtrack(nums, result)
			result = result[:len(result)-1]
		}
	}
	backtrack(nums, []int{})
	return results
}

func contains(nums []int, target int) bool {
	for _, v := range nums {
		if target == v {
			return true
		}
	}
	return false
}

func factorial(input int) (result int) {
	result = 1
	for i := input; i >= 1; i-- {
		result *= i
	}
	return
}
