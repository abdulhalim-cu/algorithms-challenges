package main

import (
	"fmt"
	"strconv"
)

func shortestPath(edges [][]string, nodeA, nodeB string) int {
	graph := buildGraphFromEdge(edges)
	return shortestPathOfGraph(graph, nodeA, nodeB, [][]string{}, map[string]bool{})
}

func buildGraphFromEdge(edges [][]string) map[string][]string {
	var graph = make(map[string][]string)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if _, ok := graph[a]; !ok {
			graph[a] = []string{}
		}
		if _, ok := graph[b]; !ok {
			graph[b] = []string{}
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	return graph
}

func shortestPathOfGraph(graph map[string][]string, src, dest string, queue [][]string, visited map[string]bool) int {
	visited[src] = true
	queue = append(queue, []string{src, "0"})
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		val, _ := strconv.Atoi(front[1])
		if front[0] == dest {
			return val
		}
		for _, neighbor := range graph[front[0]] {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = true
				queue = append(queue, []string{neighbor, strconv.Itoa(val + 1)})
			}
		}
	}
	return -1
}

func main() {
	edges := [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}
	fmt.Println(shortestPath(edges, "w", "z"))
	fmt.Println(shortestPath(edges, "y", "x"))

	edges = [][]string{
		{"a", "c"},
		{"a", "b"},
		{"c", "b"},
		{"c", "d"},
		{"b", "d"},
		{"e", "d"},
		{"g", "f"},
	}

	fmt.Println(shortestPath(edges, "a", "e"))
	fmt.Println(shortestPath(edges, "e", "c"))
	fmt.Println(shortestPath(edges, "b", "g"))

	edges = [][]string{
		{"c", "n"},
		{"c", "e"},
		{"c", "s"},
		{"c", "w"},
		{"w", "e"},
	}

	fmt.Println(shortestPath(edges, "w", "e"))
}
