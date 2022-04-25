package main

import "fmt"

func canSum(targetSum int64, numbers []int64, memo map[int64]bool) bool {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, n := range numbers {
		reminder := targetSum - n
		if canSum(reminder, numbers, memo) == true {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}

func main() {
	var memo = make(map[int64]bool)
	//fmt.Println(canSum(7, []int64{2, 3}, memo))
	//fmt.Println(canSum(7, []int64{5, 3, 4, 7}, memo))
	//fmt.Println(canSum(7, []int64{2, 4}, memo))
	//fmt.Println(canSum(8, []int64{2, 3, 5}, memo))
	fmt.Println(canSum(1000, []int64{7, 34, 54, 334, 42, 33, 14}, memo))
}
