package main

import "fmt"

func main() {
	fmt.Println('0', " ", '1')
	//fmt.Println(myAtoi("+523"))
	fmt.Println(myAtoi("18446744073709551617"))
}

func myAtoi(s string) int {
	type number struct {
		value    uint64
		negative byte
	}
	var input number
	for _, v := range s {
		if v == '-' && input.value == 0 && input.negative == 0 {
			input.negative = '-'
		} else if v == '+' && input.value == 0 && input.negative == 0 {
			input.negative = '+'
		} else if v == ' ' {
			if input.negative != 0 {
				break
			}
			continue
		} else if v >= '0' && v <= '9' {
			if input.negative == 0 {
				input.negative = '+'
			}
			input.value = input.value*10 + uint64(v) - 48
			if input.value > uint64(^(^uint32(0) >> 1)) {
				break
			}
		} else if input.value == 0 {
			return 0
		} else if input.value != 0 {
			break
		} else {
			break
		}
	}
	if input.negative == '-' {
		if input.value > uint64(^(^uint32(0) >> 1)) {
			return -int(^(^uint32(0) >> 1))
		}
		return -int(input.value)
	}
	if input.value > uint64(^uint32(0)>>1) {
		return int(^uint32(0) >> 1)
	}
	return int(input.value)
}
