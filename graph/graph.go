package main

import "fmt"

func depthFirstTraversal(graph map[string][]string, source string) []string {
	var resultList []string
	resultList = append(resultList, source)
	if nodes, ok := graph[source]; ok {
		for _, node := range nodes {
			res := depthFirstTraversal(graph, node)
			resultList = append(resultList, res...)
		}
	}
	return resultList
}

func breadthFirstTraversal(graph map[string][]string, source string) []string {
	queue := []string{source}
	var result []string
	for len(queue) > 0 {
		front := queue[0]
		result = append(result, front)
		queue = queue[1:]
		for _, node := range graph[front] {
			queue = append(queue, node)
		}
	}
	return result
}

func hasPath(graph map[string][]string, src, dest string) bool {
	if src == dest {
		return true
	}
	for _, node := range graph[src] {
		if hasPath(graph, node, dest) {
			return true
		}
	}

	return false
}

func main() {
	var graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	fmt.Println(depthFirstTraversal(graph, "a"))
	fmt.Println(breadthFirstTraversal(graph, "a"))
	fmt.Println(hasPath(graph, "b", "a"))
}
