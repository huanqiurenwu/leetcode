package main

import "fmt"

func main() {
	//fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	//fmt.Println(merge([][]int{{1,4},{4,6}}))
	fmt.Println(merge1([][]int{{1, 4}, {0, 2}, {3, 5}}))
	fmt.Println(relation([]int{1, 4}, []int{0, 2}))
}

//
func merge1(intervals [][]int) [][]int {
	//Todo: 对于切片的操作还有问题，存在切片越界的情况。
	for k := 0; k < len(intervals); k++ {
		for i := 0; i < len(intervals[k+1:]); i++ {
			base := intervals[k]
			v := intervals[k+i+1]
			if merged, relative := relation(base, v); relative == true {
				intervals[k+i+1] = merged
				intervals = append(intervals[:k], intervals[k+1:]...)
				k--
				break
			}
		}
	}
	return intervals
}

func relation(arg1 []int, arg2 []int) ([]int, bool) {
	if max(arg1[0], arg2[0]) <= min(arg1[1], arg2[1]) {
		return []int{min(arg1[0], arg2[0]), max(arg1[1], arg2[1])}, true
	}
	return []int{}, false
}

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}

func min(a1, a2 int) int {
	if a1 > a2 {
		return a2
	}
	return a1
}

func merge2(intervals [][]int) [][]int {
	//TODO:递归调用
	for k := 0; k < len(intervals); k++ {
		for i := 0; i < len(intervals[k+1:]); i++ {
			base := intervals[k]
			v := intervals[k+i+1]
			if merged, relative := relation(base, v); relative == true {
				intervals[k+i+1] = merged
				intervals = append(intervals[:k], intervals[k+1:]...)
			}
		}
	}
	return intervals
}
