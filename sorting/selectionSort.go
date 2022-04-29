package main

import "fmt"

func selectionSort(Arr []int) []int {
	for i := 0; i < len(Arr)-1; i++ {
		smallestArr := Arr[i]
		smallestIndex := i
		for j := i + 1; j < len(Arr); j++ {
			if Arr[j] < smallestArr {
				smallestIndex = j
				smallestArr = Arr[j]
			}
		}
		Arr[i], Arr[smallestIndex] = smallestArr, Arr[i]
	}
	return Arr
}

func main() {
	fmt.Println(selectionSort([]int{8, 1, 10, 4, 7, 2, 11, 6, 9, 3, 5}))
}
