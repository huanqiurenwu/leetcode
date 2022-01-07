package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge2(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	i := len(nums1) - 1
	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
			i--
		}
		nums1[i] = nums2[n]
		i--
		n--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	l := m + n - 1
	m--
	n--
	//让m和n分别指向两个数组最后一个元素
	for n >= 0 {
		if m < 0 || nums1[m] <= nums2[n] {
			nums1[l] = nums2[n]
			n--
			l--
		} else {
			nums1[l] = nums1[m]
			m--
			l--
		}
	}
}
