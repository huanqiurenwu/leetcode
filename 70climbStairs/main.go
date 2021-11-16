package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	//就是斐波拉契数列
	db := []int{1, 2}
	for i := 2; i < n; i++ {
		db = append(db, db[i-1]+db[i-2])
	}
	return db[n-1]
}
