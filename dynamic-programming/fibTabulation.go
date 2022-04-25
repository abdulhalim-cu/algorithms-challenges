package main

import "fmt"

func fibTabulation(n int) int {
	table := make([]int, n+1)
	table[1] = 1
	for i := 0; i < n; i++ {
		table[i+1] += table[i]
		if i+2 <= n {
			table[i+2] += table[i]
		}
	}
	return table[n]
}

func main() {
	fmt.Println(fibTabulation(8))
	fmt.Println(fibTabulation(50))
	fmt.Println(fibTabulation(100))
}
