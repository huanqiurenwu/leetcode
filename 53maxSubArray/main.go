package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	//设计DP数组，数组含义，dp[i]表示以元素[i]结束的和最大子序列，所以dp[i+1] = max(dp[i] + nums[i+1], nums[i+1])
	//即如果前面的最大和比0还小，就没有加上的价值，dp[i+1]就是元素nums[i+1]本身。
	max := nums[0]
	for k := 1; k < len(nums); k++ {
		if nums[k-1] > 0 {
			nums[k] = nums[k] + nums[k-1]
			if nums[k] > max {
				max = nums[k]
			}
		}
	}
	return max
}
