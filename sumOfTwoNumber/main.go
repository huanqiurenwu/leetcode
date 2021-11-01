package main

import (
	"fmt"
	"log"
	"net"
)

//func main(){
//	twoSum([]int{2,7,6,3, 8,1, 5,4},  9)
//}

func twoSum(nums []int, target int) {
	result := make([][]int, 0, 10)
	backtrace([]int{}, nums, &result, target)
	fmt.Println(result)
}

func backtrace(path []int, remain []int, result *[][]int, target int){
	if sum(path) == target {
		*result = append(*result, path)
		return
	}
	if len(remain) == 0 {
		return
	}
	for i,v := range remain{
		newPath := append([]int{}, path...)
		newPath = append(newPath, v)
		backtrace(newPath, remain[i+1:], result, target)
	}
}

//func zuhe(path []int, remain []int, target int, result *[][]int){
//	if
//	for i,v := range remain {
//
//	}
//}
//
//func zuheDiedai(remain []int, target int, result *[][]int){
//	for i,v := range remain{
//
//	}
//}

func sum(input []int) (result int){
	for _, v := range input {
		result += v
	}
	return result
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// 在一个新协程中处理客户端连接。
		go ClientHandler(conn)
	}
}
func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("捕获了一个恐慌：", v)
			log.Println("防止了程序崩溃")
		}
		c.Close()
	}()
	panic("未知错误") // 演示目的产生的一个恐慌
}