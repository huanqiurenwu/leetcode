package main

import "fmt"

func main() {
	s := "AB"
	numRows := 1
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	direction := 0
	results := make([]string, numRows)
	results[0] = results[0] + s[:1]
	s = s[1:]
	for s != "" {
		if direction > 0 {
			for i := 1; i < numRows && s != ""; i++ {
				results[direction-i] = results[direction-i] + s[:1]
				s = s[1:]
			}
			direction = 0
		} else {
			for i := 1; i < numRows && s != ""; i++ {
				results[i] = results[i] + s[:1]
				s = s[1:]
			}
			direction = numRows - 1
		}
	}
	result := ""
	for _, v := range results {
		result = result + v
	}
	return result
}
