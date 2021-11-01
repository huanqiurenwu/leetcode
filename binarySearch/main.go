package main

import "fmt"

func main(){
	fmt.Println(binarySearch1([]int{-1,0,3,5,9,12}, 0, 5, 2))
}

func binarySearch1(nums []int, start int, end int, target int) int{
	if start == end {
		fmt.Println("边界条件触发")
		if nums[start] == target{

			return start
		} else {
			return -1
		}
	}

	mid := start + (end - start) / 2
	switch {
	case nums[mid] == target:
		return mid
	case nums[mid] < target:
		return binarySearch1(nums, mid + 1, end, target)
	default:
		return binarySearch1(nums, start, mid - 1, target)
	}
}

func binarySearch2(nums []int) int{

}