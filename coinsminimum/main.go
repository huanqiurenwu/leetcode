package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(coinChange2([...]int{1,2,5}, 10))
}

//func coinChange1(coins [3]int, amount int) int{
//	var dp []int = make([]int, amount+1, amount+1)
//	for k,_ := range dp {
//		for coinIndex, coin := range coins {
//			if k - coin < 0 {
//				continue
//			}
//			dp[k] = math.MinInt(dp[k], 1 + dp(k - coin))
//		}
//	}
//}

func coinChange2(coins [3]int, amount int) int {
	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxInt
		for _, coin := range coins {
			subproblem := dp(n - coin)
			if subproblem < 0 {
				continue
			}
			res = int(math.Min(float64(res), 1 + float64(subproblem)))
		}
		if res != math.MaxInt{
			return res
		}else {
			return -1
		}
	}
	return dp(amount)
}

