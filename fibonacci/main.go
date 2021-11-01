package main

import (
	"fmt"
	"time"
)

var FIN int = 50

func main(){
	start := time.Now()
	v1 := fibonacci(FIN)
	duration1 := time.Since(start)
	fmt.Println("result:", v1, " time:", duration1)
	v2 := memofibonacci(FIN)
	duration2 := time.Since(start) - duration1
	fmt.Println("result:", v2, " time:", duration2)
}

func Triple(n int) (r int) {
	defer func() {
		r = n // 修改返回值
	}()
	return n + n
}

func fib(n int) int{
	if n <= 2 {
		return 1
	}
	if n <1{
		return 0
	}
	return fib(n-1) + fib(n-2)
}

//func memofib(n int) int{
//	var memo []int
//	memo.
//}

func fibonacci(n int) int {
	if n == 1 || n==2{
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func memofibonacci(n int) int {
	var memo map[int]int = make(map[int]int)
	return backtrack(n, memo)
}

func backtrack(n int, memo map[int]int) (result int) {
	if n == 1 || n == 2 {
		return 1
	}
	if v,present := memo[n]; present == false {
		result = backtrack(n-1, memo) + backtrack(n-2, memo)
		memo[n] = result
		return result
	} else {
		return v
	}
}