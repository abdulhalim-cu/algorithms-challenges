package main

import "fmt"

func bestSumTabulation(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)
	table[0] = []int{}
	for i := 0; i < targetSum; i++ {
		if table[i] != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					currList := table[i]
					currList = append(currList, num)
					if table[i+num] == nil || len(currList) <= len(table[i+num]) {
						table[i+num] = currList
					}
				}
			}
		}
	}
	return table[targetSum]
}

func main() {
	fmt.Println(bestSumTabulation(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSumTabulation(7, []int{2, 3, 5}))
	fmt.Println(bestSumTabulation(7, []int{5, 3, 4}))
	fmt.Println(bestSumTabulation(7, []int{2, 4}))
	fmt.Println(bestSumTabulation(8, []int{2, 3, 5}))
	fmt.Println(bestSumTabulation(8, []int{1, 4, 5}))
	fmt.Println(bestSumTabulation(100, []int{1, 2, 5, 25}))
}
