package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 19, 9}))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	var maxDp int
	dpTable := make([]int, len(nums))
	dpTable[0] = nums[0]
	dpTable[1] = nums[1]
	dpTable[2] = nums[0] + nums[2]
	maxDp = max(dpTable[1], dpTable[2])
	for i := 3; i < len(nums); i++ {
		dpTable[i] = max(dpTable[i-2]+nums[i], dpTable[i-3]+nums[i])
		if dpTable[i] > maxDp {
			maxDp = dpTable[i]
		}
	}
	return maxDp
}

func max(arr1, arr2 int) int {
	if arr1 < arr2 {
		return arr2
	}
	return arr1
}
