package main

import "fmt"

func main() {
	fmt.Println(^(^uint32(0) >> 1))
	fmt.Println((^uint32(0) >> 1))
}
