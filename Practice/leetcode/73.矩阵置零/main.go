package main

// https://leetcode.cn/problems/set-matrix-zeroes/

func setZeroes(matrix [][]int) {
	// 标记每行每列是否需要变0
	row := make([]bool, len(matrix))
	column := make([]bool, len(matrix[0]))
	// 填写标记
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 0 {
				// 标记
				row[y] = true
				column[x] = true
			}
		}
	}
	// 根据标记来遍历和修改
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if row[y] || column[x] {
				matrix[y][x] = 0
			}
		}
	}
}
