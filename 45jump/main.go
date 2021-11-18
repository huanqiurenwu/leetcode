package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 0, 4}))
}

func jump(nums []int) int {
	dpTab := make([]int, len(nums))
	for k, v := range nums {
		if k == len(nums)-1 {
			break
		}
		for i := k + 1; i <= k+v && i < len(nums); i++ {
			if dpTab[i] == 0 {
				dpTab[i] = dpTab[k] + 1
			} else {
				if dpTab[i] > dpTab[k]+1 {
					dpTab[i] = dpTab[k] + 1
				}
			}
		}
	}
	return dpTab[len(dpTab)-1]
}
