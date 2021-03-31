package main

import "fmt"

func main() {
	input := []int{1, 2, 2, 2}
	result := subsetsWithDup(input)
	fmt.Println(result)
}

func subsetsWithDup(nums []int) [][]int {
	notRepeated := make([]int, 0, 10)
	repeated := make([]int, 0, 10)
	counts := make(map[[2]int]int, 10)
	valueCount := make(map[int]int, 0)
	for _, v := range nums {
		if isExist(notRepeated, v) {
			valueCount[v]++
			counts[[2]int{len(repeated), v}] = valueCount[v]
			repeated = append(repeated, v)
		} else {
			notRepeated = append(notRepeated, v)
		}
	}
	empty := make(map[int]int, 1)
	subs := make([]map[int]int, 0, 10)
	subs = append(subs, empty)
	for _, v := range notRepeated {
		for _, sv := range subs {
			tmp := clone(sv)
			tmp[v] = 1
			subs = append(subs, tmp)
		}
	}
	for k, v := range repeated {
		for _, sv := range subs {
			if sv[v] == counts[[2]int{k, v}] {
				tmp := clone(sv)
				tmp[v]++
				subs = append(subs, tmp)
			}
		}
	}
	result := make([][]int, 0, 30)
	for _, tmpMap := range subs {
		tmpresult := make([]int, 0)
		for k, v := range tmpMap {

			for i := v; i > 0; i-- {
				tmpresult = append(tmpresult, k)
			}
		}
		result = append(result, tmpresult)
	}

	return result

}

func isExist(save []int, value int) bool {
	for _, v := range save {
		if v == value {
			return true
		}
	}
	return false
}

func clone(src map[int]int) (dst map[int]int) {
	dst = make(map[int]int, 0)
	for k, v := range src {
		dst[k] = v
	}
	return
}
