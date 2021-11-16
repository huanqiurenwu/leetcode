package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(6))
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}
