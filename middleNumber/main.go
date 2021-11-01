package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1,2}
	arr2 := []int{3,4}
	result := findMedianSortedArrays(arr1, arr2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	index1,index2 := 0,0
	var cache []int
	for index1 + index2 < len(nums1) + len(nums2) / 2{
		if nums1[index1] < nums2[index2] || index2 >= len(nums2) {
			if (index1 + index2) >= (len(nums1) + len(nums2))/ 2 - 1{
				cache = append(cache, nums1[index1])
			}
			index1 ++
		} else {
			if (index1 + index2) >= (len(nums1) + len(nums2)) / 2 - 1{
				cache = append(cache, nums2[index2])
			}
			index2 ++
		}
	}
	if  (len(nums1) + len(nums2)) % 2 == 0 {
		return float64((cache[0]+cache[1]) / 2)
	} else {
		return float64(cache[1])
	}
}

