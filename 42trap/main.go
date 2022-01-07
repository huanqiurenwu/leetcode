package main

import "fmt"

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap([]int{1, 0, 1}))
	fmt.Println(trap(input))
}

func trap(input []int) int {
	lp, rp := 0, len(input)-1
	lm, rm := input[lp], input[rp]
	result := 0
	for lp <= rp {
		if input[lp] < input[rp] {
			if lm > input[lp] {
				result += lm - input[lp]
			} else {
				lm = input[lp]
			}
			lp++
		} else {
			if rm > input[rp] {
				result += rm - input[rp]
			} else {
				rm = input[rp]
			}
			rp--
		}
	}
	return result
}
