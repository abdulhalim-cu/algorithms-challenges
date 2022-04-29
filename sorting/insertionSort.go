package main

import "fmt"

func insertionSort(Arr []int) []int {
	n := len(Arr)
	for i := 1; i < n; i++ {
		key := Arr[i]
		j := i - 1
		for j >= 0 && Arr[j] > key {
			Arr[j+1] = Arr[j]
			j -= 1
		}
		Arr[j+1] = key
	}
	return Arr
}

func insertionSortInDecreasingOrder(Arr []int) []int {
	n := len(Arr)
	for i := 1; i < n; i++ {
		key := Arr[i]
		j := i - 1
		for j >= 0 && Arr[j] < key {
			Arr[j+1] = Arr[j]
			j--
		}
		Arr[j+1] = key
	}
	return Arr
}

func main() {
	fmt.Println(insertionSort([]int{8, 1, 10, 4, 7, 2, 11, 6, 9, 3, 5}))
	fmt.Println(insertionSortInDecreasingOrder([]int{8, 1, 10, 4, 7, 2, 11, 6, 9, 3, 5}))
}
