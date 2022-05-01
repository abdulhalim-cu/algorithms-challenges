package main

import "fmt"

func largestComponent(graph map[string][]string) int {
	visited := make(map[string]bool)
	size := 0
	for node := range graph {
		currSize := exploreSize(graph, node, visited)
		if currSize > size {
			size = currSize
		}
	}
	return size
}

func exploreSize(graph map[string][]string, node string, visited map[string]bool) int {
	if _, ok := visited[node]; ok {
		return 0
	}
	visited[node] = true
	size := 1
	for _, neighbor := range graph[node] {
		size += exploreSize(graph, neighbor, visited)
	}
	return size
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
	fmt.Println(largestComponent(graph))

	graph2 := map[string][]string{
		"1": {"2"},
		"2": {"1", "8"},
		"6": {"7"},
		"9": {"8"},
		"7": {"6", "8"},
		"8": {"9", "7", "2"},
	}

	fmt.Println(largestComponent(graph2))

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
	fmt.Println(largestComponent(graph3))

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
	fmt.Println(largestComponent(graph4))

	fmt.Println(largestComponent(map[string][]string{}))
}
