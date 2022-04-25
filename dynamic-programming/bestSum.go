package main

import "fmt"

func bestSum(targetSum int64, numbers []int64, memo map[int64][]int64) []int64 {
	if v, ok := memo[targetSum]; ok {
		return v
	}
	if targetSum == 0 {
		return []int64{}
	}
	if targetSum < 0 {
		return nil
	}
	var shortestCombination []int64 = nil
	for _, num := range numbers {
		reminder := targetSum - num
		combination := bestSum(reminder, numbers, memo)
		if combination != nil {
			combination = append(combination, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}
	memo[targetSum] = shortestCombination
	return shortestCombination
}

func main() {
	var memo = make(map[int64][]int64)
	//fmt.Println(bestSum(7, []int64{5, 3, 4, 7}, memo))
	//fmt.Println(bestSum(8, []int64{2, 3, 5}, memo))
	//fmt.Println(bestSum(8, []int64{1, 4, 5}, memo))
	//fmt.Println(bestSum(100, []int64{1, 2, 5, 25}, memo))
	fmt.Println(bestSum(1000, []int64{35, 40, 45, 49, 55, 58, 60, 110, 135, 145, 255, 350, 545}, memo))
}
