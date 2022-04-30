package main

import "fmt"

func mergeSort(Arr []int) []int {
	if len(Arr) < 2 {
		return Arr
	}
	left := mergeSort(Arr[:len(Arr)/2])
	right := mergeSort(Arr[len(Arr)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var final []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}

	for i < len(left) {
		final = append(final, left[i])
		i++
	}

	for j < len(right) {
		final = append(final, right[j])
		j++
	}
	return final
}

func main() {
	Arr := []int{8, 1, 10, 4, 7, 2, 11, 6, 9, 3, 5}

	fmt.Println(mergeSort(Arr))
}
