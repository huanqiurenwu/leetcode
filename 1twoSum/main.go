package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 7))
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	for k1, v1 := range nums[:len(nums)-1] {
		for k2, v2 := range nums[k1+1:] {
			if v1+v2 == target {
				result = append(result, k1, k2+k1+1)
				return result
			}
		}
	}
	return result
}
