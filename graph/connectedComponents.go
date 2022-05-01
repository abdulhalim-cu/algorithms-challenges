package main

import "fmt"

func connectedComponentsCount(graph map[string][]string) int {
	var visited = make(map[string]bool)
	count := 0

	for node := range graph {
		if explore(graph, node, visited) {
			count++
		}
	}
	return count
}

func explore(graph map[string][]string, node string, visited map[string]bool) bool {
	if _, ok := visited[node]; ok {
		return false
	}
	visited[node] = true
	for _, neighbor := range graph[node] {
		explore(graph, neighbor, visited)
	}
	return true
}

func main() {
	graph := map[string][]string{
		"0": {"8", "1", "5"},
		"1": {"0"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"2": {"3", "4"},
		"3": {"2", "4"},
		"4": {"3", "2"},
	}
	fmt.Println(connectedComponentsCount(graph))

	graph2 := map[string][]string{
		"1": {"2"},
		"2": {"1", "8"},
		"6": {"7"},
		"9": {"8"},
		"7": {"6", "8"},
		"8": {"9", "7", "2"},
	}
	fmt.Println(connectedComponentsCount(graph2))

	graph3 := map[string][]string{
		"3": {},
		"4": {"6"},
		"6": {"4", "5", "7", "8"},
		"8": {"6"},
		"7": {"6"},
		"5": {"6"},
		"1": {"2"},
		"2": {"1"},
	}

	fmt.Println(connectedComponentsCount(graph3))

	graph4 := map[string][]string{
		"0": {"4", "7"},
		"1": {},
		"2": {},
		"3": {"6"},
		"4": {"0"},
		"6": {"3"},
		"7": {"0"},
		"8": {},
	}
	fmt.Println(connectedComponentsCount(graph4))
}
