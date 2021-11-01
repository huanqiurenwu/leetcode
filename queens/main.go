package main

import "fmt"

func main(){
	fmt.Println(solveNQueens())
	fmt.Println('.')
}

//func solveNQueens(n int){
//	var board [8][8]int
//	var backtrack func(*[8][8]int, int)
//	result := [][8][8]int{}
//	backtrack = func(board *[8][8]int, row int){
//		if row == 8 {
//			result = append(result, *board)
//			return
//		}
//		for
//	}
//	backtrack(&board, 0)
//}
//
//func isValid(board *[8][2]int, row int, col int) bool{
//	for i := 0; i < row; i++ {
//		if board[i][col] == 1 {
//			return false
//		}
//	}
//}

func solveNQueens()[][8]int{
	var axisOfQueens [][8]int
	var backtrack func(*[8]int, int)bool
	axises := [8]int{}
	backtrack = func(axises *[8]int, row int)bool{
		if row == 8 {
			axisOfQueens = append(axisOfQueens, *axises)
			return true
		}
		for i := 0; i<8;i++{
			if isValid(axises, row, i){
				axises[row] = i
				if backtrack(axises, row+1) {
					return true
				}
				axises[row] = 0
			}
		}
		return false
	}
	backtrack(&axises, 0)
	return axisOfQueens
}

func isValid(axises *[8]int, x int, y int) bool {
	for i := 0; i<x;i++{
		if x - i == y - axises[i] || x - i == -(y - axises[i]) || y == axises[i]{
			return false
		}
	}
	return true
}
