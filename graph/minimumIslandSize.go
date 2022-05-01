package main

import (
	"fmt"
	"math"
	"strconv"
)

func minIslandCount(grid [][]string) int {
	var visited = make(map[string]bool, len(grid)+len(grid[0]))
	if grid == nil {
		return 0
	}
	var minSize int = math.MaxInt
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			size := exploreIslandSize(grid, r, c, visited)
			if size > 0 && size < minSize {
				minSize = size
			}
		}
	}
	return minSize
}

func exploreIslandSize(grid [][]string, r, c int, visited map[string]bool) int {
	roWInbounds := 0 <= r && r < len(grid)
	colInbounds := 0 <= c && c < len(grid[0])
	if !roWInbounds || !colInbounds {
		return 0
	}
	if grid[r][c] == "W" {
		return 0
	}
	pos := strconv.Itoa(r) + "," + strconv.Itoa(c)
	if _, ok := visited[pos]; ok {
		return 0
	}
	visited[pos] = true

	size := 1
	size += exploreIslandSize(grid, r-1, c, visited)
	size += exploreIslandSize(grid, r+1, c, visited)
	size += exploreIslandSize(grid, r, c-1, visited)
	size += exploreIslandSize(grid, r, c+1, visited)
	return size
}

func main() {
	grid := [][]string{
		{"W", "L", "W", "W", "W"},
		{"W", "L", "W", "W", "W"},
		{"W", "W", "W", "L", "W"},
		{"W", "W", "L", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"L", "L", "W", "W", "W"},
	}
	fmt.Println(minIslandCount(grid))

	grid = [][]string{
		{"L", "W", "W", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"W", "L", "W", "L", "W"},
		{"W", "W", "W", "W", "W"},
		{"W", "W", "L", "L", "L"},
	}
	fmt.Println(minIslandCount(grid))

	grid = [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}
	fmt.Println(minIslandCount(grid))

	grid = [][]string{
		{"W", "W"},
		{"W", "W"},
		{"W", "W"},
	}
	fmt.Println(minIslandCount(grid))
}
