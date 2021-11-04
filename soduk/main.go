package main

import "fmt"

func main() {
	input := [][]byte{{5, 3, '.', '.', 7, '.', '.', '.', '.'}, {6, '.', '.', 1, 9, 5, '.', '.', '.'}, {'.', 9, 8, '.', '.', '.', '.', 6, '.'}, {8, '.', '.', '.', 6, '.', '.', '.', 3}, {4, '.', '.', 8, '.', 3, '.', '.', 1}, {7, '.', '.', '.', 2, '.', '.', '.', 6}, {'.', 6, '.', '.', '.', '.', 2, 8, '.'}, {'.', '.', '.', 4, 1, 9, '.', '.', 5}, {'.', '.', '.', '.', 8, '.', '.', 7, 9}}
	printSudoku(input)
	result := solveSudoku(input)
	for _, v := range result {
		printSudoku(v)
	}
}

func printSudoku(board [][]byte) {
	//@TODO: 结果计算不出来。需要改进
	for _, cols := range board {
		for _, v := range cols {
			if v == '.' {
				fmt.Printf("%c ", v)
				continue
			}
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}

}

func solveSudoku(board [][]byte) [][][]byte {
	var backtrack func([][]byte, int)
	result := [][][]byte{}
	backtrack = func(board [][]byte, row int) {
		if row == 9 {
			result = append(result, board)
			return
		}
		for i := 0; i < 9; i++ {
			for v := 1; v <= 9; v++ {
				if isValid(row, i, byte(v), board) {
					board[row][i] = byte(v)
					backtrack(board, row+1)
					board[row][i] = '.'
				}
			}
		}
	}
	backtrack(board, 0)
	return result
}

func isValid(x, y int, value byte, board [][]byte) bool {
	if board[x][y] != '.' {
		return false
	}
	//判断每行，
	for i := 0; i < 9 && i != y; i++ {
		if board[x][i] == value {
			return false
		}
	}
	//判断每列
	for i := 0; i < 9 && i != x; i++ {
		if board[i][y] == value {
			return false
		}
	}
	//判断每个宫
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if x/3*3+i == x || y/3*3+j == y {
				continue
			}
			if board[x/3*3+i][y/3*3+j] == value {
				return false
			}
		}
	}
	return true
}
