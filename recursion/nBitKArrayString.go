package main

import "fmt"

func kArrayString(Arr []int, n, k int, pos int) {
	if n < 1 {
		fmt.Println(Arr)
		return
	}
	for i := 0; i < k; i++ {
		Arr[pos] = i
		kArrayString(Arr, n-1, k, pos+1)
	}
}

func main() {
	Arr := make([]int, 4)
	kArrayString(Arr, 4, 2, 0)
}
