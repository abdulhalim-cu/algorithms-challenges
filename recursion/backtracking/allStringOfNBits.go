package main

import "fmt"

func nBitBinaryString(A []int, n int, pos int) {
	if n < 1 {
		fmt.Println(A)
		return
	} else {
		A[pos] = 0
		nBitBinaryString(A, n-1, pos+1)
		A[pos] = 1
		nBitBinaryString(A, n-1, pos+1)
	}
}

func main() {
	Arr := make([]int, 3)
	nBitBinaryString(Arr, 3, 0)
}
