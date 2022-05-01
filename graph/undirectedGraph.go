package main

import "fmt"

func undirectedPath(edges [][]string, nodeA, nodeB string) bool {
	graph := buildGraph(edges)
	return hasUndirectedPath(graph, nodeA, nodeB, []string{})
}

func hasUndirectedPath(graph map[string][]string, src, dest string, visited []string) bool {
	if src == dest {
		return true
	}
	if contains(visited, src) {
		return false
	}

	visited = append(visited, src)

	for _, node := range graph[src] {
		if hasUndirectedPath(graph, node, dest, visited) {
			return true
		}
	}
	return false
}

func contains(slice []string, s string) bool {
	for i := range slice {
		if slice[i] == s {
			return true
		}
	}
	return false
}

func buildGraph(edges [][]string) map[string][]string {
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

func main() {
	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}

	res := undirectedPath(edges, "i", "l")
	fmt.Println(res)
	res = undirectedPath(edges, "j", "m")
	fmt.Println(res)
	res = undirectedPath(edges, "m", "i")
	fmt.Println(res)
	res = undirectedPath(edges, "n", "o")
	fmt.Println(res)
	res = undirectedPath(edges, "n", "k")
	fmt.Println(res)
	res = undirectedPath(edges, "i", "o")
	fmt.Println(res)
	res = undirectedPath(edges, "o", "n")
	fmt.Println(res)
}
