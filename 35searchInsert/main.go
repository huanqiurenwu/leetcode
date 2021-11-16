package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{0, 1, 2, 3}, -1))
}
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return left
}
