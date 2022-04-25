package main

import (
	"fmt"
	"strconv"
)

func gridTraveler(n, m int, memo map[string]int64) int64 {
	strN := strconv.Itoa(n)
	strM := strconv.Itoa(m)
	key := strN + "," + strM
	if v, ok := memo[key]; ok {
		return v
	}
	if n == 0 || m == 0 {
		memo[key] = 0
		return 0
	}
	if n == 1 && m == 1 {
		memo[key] = 1
		return 1
	}
	memo[key] = gridTraveler(n-1, m, memo) + gridTraveler(n, m-1, memo)
	return memo[key]
}

func main() {
	var memo = make(map[string]int64)
	m := gridTraveler(50, 50, memo) //858110510779117752
	fmt.Println(m)
}
