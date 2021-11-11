package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("ababc"))
}

func longestPalindrome(s string) string {
	bytes := []byte(s)
	dpMap := make(map[int]string, len(s)) // key是长度，value是回文子串。
	//第1步选择作为结束字符的字符
	maxLen := 0
	for k, _ := range bytes {
		for i, _ := range bytes[:k+1] {
			if isPalindrome(bytes[i : k+1]) {
				if len(bytes[i:k+1]) > maxLen {
					maxLen = len(bytes[i : k+1])
				}
				dpMap[len(bytes[i:k+1])] = string(bytes[i : k+1])
			}
		}
	}
	return dpMap[maxLen]
}

func isPalindrome(input []byte) (result bool) {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
