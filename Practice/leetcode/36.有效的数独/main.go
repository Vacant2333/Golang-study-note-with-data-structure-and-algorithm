package main

// https://leetcode.cn/problems/valid-sudoku/
/*
[[".","8","7","6","5","4","3","2","1"],
["2",".",".",".",".",".",".",".","."],
["3",".",".",".",".",".",".",".","."],
["4",".",".",".",".",".",".",".","."],
["5",".",".",".",".",".",".",".","."],
["6",".",".",".",".",".",".",".","."],
["7",".",".",".",".",".",".",".","."],
["8",".",".",".",".",".",".",".","."],
["9",".",".",".",".",".",".",".","."]]
*/
func isValidSudoku(board [][]byte) bool {
	// 行
	row := make([]map[byte]bool, 9)
	// 列
	column := make([]map[byte]bool, 9)
	// 九宫格 0~8
	ward := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		column[i] = make(map[byte]bool)
		ward[i] = make(map[byte]bool)
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			v := board[y][x]
			if v == '.' {
				continue
			}
			if _, ok := row[y][v]; ok {
				return false
			} else {
				row[y][v] = true
			}
			if _, ok := column[x][v]; ok {
				return false
			} else {
				column[x][v] = true
			}
			// 九宫格的下标,y要先/3再*3
			wardIndex := x/3 + (y/3)*3
			if _, ok := ward[wardIndex][v]; ok {
				return false
			} else {
				ward[wardIndex][v] = true
			}
		}
	}
	return true
}
