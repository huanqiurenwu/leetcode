package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(123456789999999987654321))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		reverse := x % 10
		tmpx := x / 10
		for ; tmpx > 0; reverse, tmpx = reverse*10+tmpx%10, tmpx/10 {

		}
		if reverse == x {
			return true
		}
	}
	return false
}
