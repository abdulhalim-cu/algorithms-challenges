package main

import "fmt"

func fib(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func main() {
	var memo = make(map[int]int)
	num := fib(50, memo)
	fmt.Println(num)
}
