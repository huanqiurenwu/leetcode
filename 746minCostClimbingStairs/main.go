package main

import "fmt"

func main() {
	//fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairsDPVersion([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairsDPVersion([]int{1, 100, 1}))
}

func minCostClimbingStairsRecursion(cost []int) int {
	return min(minCostClimbingStairsHelper(cost), minCostClimbingStairsHelper(cost[:len(cost)-1]))
}

func minCostClimbingStairsHelper(cost []int) int {
	if len(cost) == 2 {
		return cost[1]
	}
	if len(cost) == 1 {
		return cost[0]
	}
	return min(minCostClimbingStairsHelper(cost[:len(cost)-1])+cost[len(cost)-1], minCostClimbingStairsHelper(cost[:len(cost)-2])+cost[len(cost)-1])
}
func min(arr1, arr2 int) int {
	if arr1 < arr2 {
		return arr1
	}
	return arr2
}

func minCostClimbingStairsDPVersion(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}
	dpTable := make([]int, len(cost))
	dpTable[0] = cost[0]
	dpTable[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dpTable[i] = min(dpTable[i-1]+cost[i], dpTable[i-2]+cost[i])
	}
	return min(dpTable[len(cost)-1], dpTable[len(cost)-2])
}
