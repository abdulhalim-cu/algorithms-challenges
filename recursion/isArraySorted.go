package main

import "fmt"

func isArraySorted(Arr []int, n int) bool {
	if n == 1 {
		return true
	}
	if Arr[n-1] < Arr[n-2] {
		return false
	}
	return isArraySorted(Arr, n-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(isArraySorted(arr, len(arr)))
	arr2 := []int{1, 3, 2, 5, 4, 6, 7, 8}
	fmt.Println(isArraySorted(arr2, len(arr2)))
}
