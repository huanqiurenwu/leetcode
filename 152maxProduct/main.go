package main

import "fmt"

func main() {
	maxProduct([]int{2, 3, -2, 4, -1})
	//正负数组两个属性：
	//0的含义，当前值只能为正值或者为负值； 如果正负数组里面都有值，那么
	//正数：[2  6   0   4  48]
	//负数：[0  0 -12 -48  -4]

	maxProduct([]int{-2, 0, 1})
	maxProduct([]int{-1, -2, 2})
}

func maxProduct(nums []int) int {
	positive := make([]int, len(nums))
	negative := make([]int, len(nums))
	if nums[0] > 0 {
		positive[0] = nums[0]
	} else if nums[0] < 0 {
		negative[0] = nums[0]
	}
	for k := 1; k < len(nums); k++ {
		if nums[k] > 0 {
			if nums[k]*positive[k-1] > nums[k] {
				positive[k] = nums[k] * positive[k-1]
			} else {
				positive[k] = nums[k]
			}
			negative[k] = nums[k] * negative[k-1]
		} else if nums[k] < 0 {
			if nums[k]*positive[k-1] < nums[k] {
				negative[k] = positive[k-1] * nums[k]
			} else {
				negative[k] = nums[k]
			}
			positive[k] = negative[k-1] * nums[k]
		}
	}
	fmt.Println(negative)
	fmt.Println(positive)
	maxInt := nums[0]
	for k, v := range positive {
		if v > 0 {
			if maxInt < v {
				maxInt = v
			}
		} else {
			if maxInt < negative[k] {
				maxInt = negative[k]
			}
		}
	}
	return maxInt

	//for k := 1; k < len(nums); k++ {
	//	if nums[k] == 0 {
	//		continue
	//	}
	//	if nums[k] > 0 {
	//		if nums[k] >= 1{
	//
	//		} else{
	//			positive[k] = nums[k]
	//			negative[k] = 1
	//		}
	//	} else if nums[k] < 0 {
	//		if nums[k] <= -1 {
	//
	//		} else {
	//
	//		}
	//	}
	//}
}
