package main

import "fmt"

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fastSort(nums)
	fmt.Println(nums)
	result := fourSum(nums, 0)
	fmt.Println(result)
}

func fourSum(nums []int, target int) (result [][]int) {
	fastSort(nums)
	for i := 0; i < len(nums) && nums[i] < target; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum := target - nums[i]
		for j := i + 1; j < len(nums) && nums[j] < threeSum; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			twoSum := threeSum - nums[j]
			for le, ri := j+1, len(nums)-1; le < ri; {
				if ri < len(nums)-1 && nums[ri] == nums[ri+1] {
					ri--
					continue
				}
				if le > j+1 && nums[le] == nums[le-1] {
					le++
					continue
				}
				if nums[le] < twoSum-nums[ri] {
					le++
				} else if twoSum-nums[ri] < nums[le] {
					ri--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[le], nums[ri]})
					break
				}
			}
		}
	}
	return
}

func fastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := 0
	le, ri := mid+1, len(nums)-1
	for le <= ri {
		for ; le <= ri && nums[le] <= nums[mid]; le++ {
		}
		for ; le <= ri && nums[ri] > nums[mid]; ri-- {
		}

		if le < ri {
			nums[le], nums[ri] = nums[ri], nums[le]
		}
	}
	nums[mid], nums[ri] = nums[ri], nums[mid]
	fastSort(nums[:ri])
	fastSort(nums[le:])
}
