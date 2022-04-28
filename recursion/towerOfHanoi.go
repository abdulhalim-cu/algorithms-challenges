package main

import "fmt"

func towersOfHanoi(n int, source string, dest string, aux string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, dest)
		return
	}
	towersOfHanoi(n-1, source, aux, dest)
	fmt.Printf("Move disk %d from %s to %s\n", n, source, dest)
	towersOfHanoi(n-1, aux, dest, source)
}

func main() {
	towersOfHanoi(3, "Src", "Dest", "Aux")
}
