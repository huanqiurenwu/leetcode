package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for ; i <= j; i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			i--
			j--
		}
	}
	return i
}
