package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	currentLongestSubs := []byte{}
	bytes := []byte(s)
	maxLen := 1
	for _, v := range bytes {
		if index, contained := contains(currentLongestSubs, v); contained {
			currentLongestSubs = append(currentLongestSubs[index+1:], v)
		} else {
			currentLongestSubs = append(currentLongestSubs, v)
		}
		//printBytes(currentLongestSubs)
		if maxLen < len(currentLongestSubs) {
			maxLen = len(currentLongestSubs)
		}
	}
	return maxLen

}

func contains(container []byte, value byte) (int, bool) {
	for k, v := range container {
		if value == v {
			return k, true
		}
	}
	return 0, false
}

func printBytes(input []byte) {
	for _, v := range input {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
