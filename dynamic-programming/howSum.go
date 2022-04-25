package main

import (
	"fmt"
)

func howSum(targetSum int64, numbers []int64, memo map[int64][]int64) []int64 {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 {
		return []int64{}
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		reminder := targetSum - num
		result := howSum(reminder, numbers, memo)
		if result != nil {
			result = append(result, num)
			memo[targetSum] = result
			return memo[targetSum]
		}
	}
	memo[targetSum] = nil
	return nil
}

func main() {
	var memo = make(map[int64][]int64)
	fmt.Println(howSum(300, []int64{7, 14}, memo))
}
