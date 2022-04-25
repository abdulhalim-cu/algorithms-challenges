package main

import "fmt"

func howSumTabulation(targetSum int, numbers []int) []int {
	table := make([][]int, targetSum+1)
	table[0] = []int{}

	for i := 0; i < targetSum; i++ {
		if table[i] != nil {
			for _, num := range numbers {
				if i+num <= targetSum {
					currVal := table[i]
					currVal = append(currVal, num)
					table[i+num] = currVal
				}
			}
		}
	}
	return table[targetSum]
}

func main() {
	fmt.Println(howSumTabulation(7, []int{2, 3}))
	fmt.Println(howSumTabulation(7, []int{5, 3, 4}))
	fmt.Println(howSumTabulation(7, []int{2, 4}))
	fmt.Println(howSumTabulation(8, []int{2, 3, 5}))
	fmt.Println(howSumTabulation(300, []int{7, 14}))
}
