package main

import (
	"fmt"
	//"sort"
)

func main() {
	nums := []int{1, 2, 3, 1}
	//sort.Ints(nums)
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[1], nums[0])
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := second
		second = max(first+nums[i], second)
		first = tmp
	}
	return second
}

func max(arr1, arr2 int) int {
	if arr1 > arr2 {
		return arr1
	}
	return arr2
}
