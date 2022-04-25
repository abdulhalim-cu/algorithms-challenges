package main

import (
	"fmt"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func depthFirstValues(n *Node) []string {
	if n == nil {
		return []string{}
	}
	// using recursion
	//leftNodeValues := depthFirstValues(n.Left)
	//rightNodeValues := depthFirstValues(n.Right)
	//result := []string{n.Val}
	//result = append(result, leftNodeValues...)
	//result = append(result, rightNodeValues...)
	//return result

	// using for loop
	var result []string
	stack := []*Node{n}
	for len(stack) > 0 {
		lastIndex := len(stack) - 1
		curr := stack[lastIndex]
		stack = stack[:lastIndex]
		result = append(result, curr.Val)

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}

		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return result
}

func breathFirstValues(n *Node) []string {
	if n == nil {
		return []string{}
	}
	var result []string
	queue := []*Node{n}
	for len(queue) > 0 {
		curr := queue[0]
		result = append(result, curr.Val)
		queue = queue[1:]
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return result
}

func main() {
	a := &Node{}
	a.Val = "a"
	b := &Node{Val: "b"}
	c := &Node{Val: "c"}
	d := &Node{Val: "d"}
	e := &Node{Val: "e"}
	f := &Node{Val: "f"}
	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Right = f

	fmt.Println(depthFirstValues(a))
	fmt.Println(breathFirstValues(a))
}
