package main

import (
	"fmt"
	"strconv"
)

func islandCount(grid [][]string) int {
	var visited = make(map[string]bool, len(grid)+len(grid[0]))
	if grid == nil {
		return 0
	}
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if exploreIsland(grid, r, c, visited) {
				count++
			}
		}
	}
	return count
}

func exploreIsland(grid [][]string, r, c int, visited map[string]bool) bool {
	roWInbounds := 0 <= r && r < len(grid)
	colInbounds := 0 <= c && c < len(grid[0])
	if !roWInbounds || !colInbounds {
		return false
	}
	if grid[r][c] == "W" {
		return false
	}
	pos := strconv.Itoa(r) + "," + strconv.Itoa(c)
	if _, ok := visited[pos]; ok {
		return false
	}
	visited[pos] = true
	exploreIsland(grid, r-1, c, visited)
	exploreIsland(grid, r+1, c, visited)
	exploreIsland(grid, r, c-1, visited)
	exploreIsland(grid, r, c+1, visited)
	return true
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
	fmt.Println(islandCount(grid))

	grid = [][]string{
		{"L", "W", "W", "L", "W"},
		{"L", "W", "W", "L", "L"},
		{"W", "L", "W", "L", "W"},
		{"W", "W", "W", "W", "W"},
		{"W", "W", "L", "L", "L"},
	}
	fmt.Println(islandCount(grid))

	grid = [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}
	fmt.Println(islandCount(grid))

	grid = [][]string{
		{"W", "W"},
		{"W", "W"},
		{"W", "W"},
	}
	fmt.Println(islandCount(grid))
}
