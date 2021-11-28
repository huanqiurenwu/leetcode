package main

import "fmt"

func main() {
	//nums := []int{8,8,8,8,5,8,8}
	nums := []int{5, 1, 2, 6, 3, 7, 8}
	//nums := []int{4,7,5}
	fastSort(nums)
	fmt.Println(nums)
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}

func permuteUnique(nums []int) [][]int {
	fastSort(nums) //先对数组进行排序
	result := [][]int{}
	var backtrack func([]int, []int)
	used := make([]bool, len(nums))
	//used用来标记递归树的纵向路径上是否被选中过。
	backtrack = func(nums []int, trace []int) {
		if len(trace) == len(nums) {
			tmp := append([]int{}, trace...)
			result = append(result, tmp)
			return
		}
		forInUsed := make([]bool, len(nums))
		//forInUsed用来在循环中标记同一深度的分支是否被选中过
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && forInUsed[i-1] == true {
				forInUsed[i] = true
				continue
			}
			forInUsed[i] = true
			trace = append(trace, nums[i])
			used[i] = true
			backtrack(nums, trace)
			used[i] = false
			trace = trace[:len(trace)-1]
		}
	}
	backtrack(nums, []int{})
	return result
}

func fastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j := 1, len(nums)-1
	mid := nums[0]
	for i <= j {
		for ; i <= j && mid > nums[i]; i++ {
			//退出条件：遇到大于等于它的元素，此时指向大于等于他的元素
		}
		for ; i <= j && mid <= nums[j]; j-- {
			//退出条件：遇到小于它的元素，此时指向小于他的元素
			//所以 两端循环执行完成的时候，i不会等于j，要么i和j还没相交，极限情况下i = j - 1；要么i和j相交过了，此时j=i-1
			//所以下面的j--,i++相差2也只会让j=i-1，而不会有其他退出条件
			//使得退出条件唯一
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			i++
		}
	}
	nums[j], nums[0] = mid, nums[j]
	fastSort(nums[0:j])
	fastSort(nums[i:])
}
