package main

import "fmt"

func gridTravelerTabulation(row, col int) int {
	table := make([][]int, row+1)
	for i := range table {
		table[i] = make([]int, col+1)
	}
	table[1][1] = 1
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			current := table[i][j]
			if j+1 <= col {
				table[i][j+1] += current
			}
			if i+1 <= row {
				table[i+1][j] += current
			}
		}
	}
	return table[row][col]
}

func main() {
	fmt.Println(gridTravelerTabulation(100, 100))
}
