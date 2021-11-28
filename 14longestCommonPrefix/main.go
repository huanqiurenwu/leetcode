package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"aaa", "ab", "aa"}))
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
Outer:
	for i := 0; ; i++ {
		if i < len(strs[0]) {
			prefix = strs[0][:i+1]
		} else {
			break
		}
		for _, v := range strs {
			if len(prefix) > len(v) { //较短的串是前缀
				prefix = prefix[0 : len(prefix)-1]
				break Outer
			}
			if prefix[i] != v[i] { //前缀要和字符串相等
				prefix = prefix[0 : len(prefix)-1]
				break Outer
			}
		}

	}
	return prefix
}
