package main

import "fmt"

func quickSort(Arr []int) []int {
	if len(Arr) < 2 {
		return Arr
	}
	pivot := Arr[len(Arr)/2]
	var leftArr []int
	var rightArr []int
	for _, ar := range Arr {
		if ar < pivot {
			leftArr = append(leftArr, ar)
		}
		if ar > pivot {
			rightArr = append(rightArr, ar)
		}
	}
	var result []int
	left := quickSort(leftArr)
	result = append(result, left...)
	result = append(result, pivot)
	right := quickSort(rightArr)
	result = append(result, right...)
	return result
}

func main() {
	fmt.Println(quickSort([]int{8, 1, 10, 4, 7, 2, 11, 6, 9, 3, 5}))
}
