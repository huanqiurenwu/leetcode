package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea2([]int{1, 1}))
	fmt.Println(maxArea2([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea2([]int{1, 2, 1}))
}

func maxArea(height []int) int {
	//双指针解法
	areaMax := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			tmpArea := height[i] * (j - i)
			if tmpArea > areaMax {
				areaMax = tmpArea
			}
			i++
		} else {
			tmpArea := height[j] * (j - i)
			if tmpArea > areaMax {
				areaMax = tmpArea
			}
			j--
		}
	}
	return areaMax
}

func maxArea2(height []int) int {
	//暴力解法
	areaMax := 0
	for k1, v1 := range height[:len(height)-1] {
		for k2, v2 := range height[k1+1:] {
			area := 0
			if v1 < v2 {
				area = (k2 + 1) * v1
			} else {
				area = (k2 + 1) * v2
			}
			if area > areaMax {
				areaMax = area
			}
		}
	}
	return areaMax
}
