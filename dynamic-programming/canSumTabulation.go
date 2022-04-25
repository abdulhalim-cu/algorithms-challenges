package main

import "fmt"

func canSumTabulation(targetSum int, numbers []int) bool {
	table := make([]bool, targetSum+1)

	table[0] = true
	for i := 0; i < targetSum; i++ {
		if table[i] == false {
			continue
		}
		for _, num := range numbers {
			if i+num <= targetSum {
				table[i+num] = true
			}
		}
	}
	return table[targetSum]
}

func main() {

	//fmt.Println(canSumTabulation(7, []int{2, 3}))
	//fmt.Println(canSumTabulation(7, []int{5, 3, 4}))
	//fmt.Println(canSumTabulation(7, []int{2, 4}))
	//fmt.Println(canSumTabulation(8, []int{5, 3, 2}))
	fmt.Println(canSumTabulation(300, []int{7, 14}))
}
