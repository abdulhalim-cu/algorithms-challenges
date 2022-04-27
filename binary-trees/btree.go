package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type NodeInt struct {
	Val   int
	Left  *NodeInt
	Right *NodeInt
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

func depthFirstValuesNodeInt(n *NodeInt) []int {
	if n == nil {
		return []int{}
	}
	var result []int
	stack := []*NodeInt{n}
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

func breathFirstValuesNodeInt(n *NodeInt) []int {
	if n == nil {
		return []int{}
	}
	var result []int
	queue := []*NodeInt{n}
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

// iterative version
/*
func treeIncludes(n *Node, targetVal string) bool {
	if n == nil {
		return false
	} else if n.Val == targetVal {
		return true
	}
	queue := []*Node{n}
	found := false
	for len(queue) > 0 {
		curr := queue[0]
		if curr.Left != nil && curr.Left.Val == targetVal {
			found = true
			break
		} else if curr.Right != nil && curr.Right.Val == targetVal {
			found = true
			break
		}
		queue = queue[1:]
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return found
}*/

// Recursive version
func treeIncludes(node *Node, targetVal string) bool {
	if node == nil {
		return false
	}
	if node.Val == targetVal {
		return true
	}
	return treeIncludes(node.Left, targetVal) || treeIncludes(node.Right, targetVal)
}

func treeIncludesNodeInt(node *NodeInt, targetVal int) bool {
	if node == nil {
		return false
	}
	if node.Val == targetVal {
		return true
	}
	return treeIncludesNodeInt(node.Left, targetVal) || treeIncludesNodeInt(node.Right, targetVal)
}

func treeSum(node *NodeInt) int {
	if node == nil {
		return 0
	}
	return node.Val + treeSum(node.Left) + treeSum(node.Right)
}

func treeMin(node *NodeInt) int {
	if node == nil {
		return math.MaxInt64
	}
	a := node.Val
	b := treeMin(node.Left)
	c := treeMin(node.Right)
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}

func maxPathSum(node *NodeInt) int {
	if node == nil {
		return 0
	}
	l := node.Val + maxPathSum(node.Left)
	r := node.Val + maxPathSum(node.Right)
	if l > r {
		return l
	}
	return r
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
	fmt.Println(treeIncludes(b, "c"))
	fmt.Println(treeIncludes(b, "e"))
	fmt.Println(treeIncludes(c, "f"))
	fmt.Println(treeIncludes(c, "g"))

	newNodeInt := &NodeInt{
		Val: 10,
		Left: &NodeInt{
			Val: 14,
			Left: &NodeInt{
				Val: 18,
				Left: &NodeInt{
					Val: 21,
					Left: &NodeInt{
						Val: 13,
						Left: &NodeInt{
							Val: 15,
						},
						Right: &NodeInt{
							Val: 19,
							Left: &NodeInt{
								Val: 22,
								Right: &NodeInt{
									Val: 23,
								},
							},
							Right: &NodeInt{
								Val: 24,
							},
						},
					},
					Right: &NodeInt{
						Val: 25,
						Right: &NodeInt{
							Val: 26,
						},
					},
				},
				Right: &NodeInt{
					Val: 26,
					Left: &NodeInt{
						Val: 27,
						Left: &NodeInt{
							Val: 28,
						},
						Right: &NodeInt{
							Val: 29,
						},
					},
					Right: &NodeInt{
						Val: 30,
						Left: &NodeInt{
							Val: 31,
						},
						Right: &NodeInt{
							Val: 32,
						},
					},
				},
			},
			Right: &NodeInt{
				Val: 33,
				Left: &NodeInt{
					Val: 34,
					Left: &NodeInt{
						Val: 35,
					},
					Right: &NodeInt{
						Val: 36,
					},
				},
				Right: &NodeInt{
					Val: 37,
				},
			},
		},
		Right: &NodeInt{
			Val: 39,
			Left: &NodeInt{
				Val: 40,
				Left: &NodeInt{
					Val: 41,
					Left: &NodeInt{
						Val: 42,
					},
					Right: &NodeInt{
						Val: 43,
						Left: &NodeInt{
							Val: 44,
						},
						Right: &NodeInt{
							Val: 45,
							Left: &NodeInt{
								Val: 46,
							},
						},
					},
				},
				Right: &NodeInt{
					Val: 47,
					Left: &NodeInt{
						Val: 48,
					},
					Right: &NodeInt{
						Val: 49,
						Left: &NodeInt{
							Val: 50,
						},
						Right: &NodeInt{
							Val: 51,
							Left: &NodeInt{
								Val: 52,
							},
							Right: &NodeInt{
								Val: 53,
								Left: &NodeInt{
									Val: 54,
								},
								Right: &NodeInt{
									Val: 55,
								},
							},
						},
					},
				},
			},
			Right: &NodeInt{
				Val: 56,
				Left: &NodeInt{
					Val: 57,
				},
				Right: &NodeInt{
					Val: 58,
					Left: &NodeInt{
						Val: 59,
					},
					Right: &NodeInt{
						Val: 60,
						Left: &NodeInt{
							Val: 61,
						},
						Right: &NodeInt{
							Val: 62,
							Left: &NodeInt{
								Val: 63,
							},
							Right: &NodeInt{
								Val: 64,
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(depthFirstValuesNodeInt(newNodeInt))
	fmt.Println(breathFirstValuesNodeInt(newNodeInt))
	fmt.Println(treeIncludesNodeInt(newNodeInt, 60))
	fmt.Println(treeIncludesNodeInt(newNodeInt, 63))
	fmt.Println(treeIncludesNodeInt(newNodeInt, 65))
	fmt.Println(treeIncludesNodeInt(newNodeInt, 66))

	lst := "10 14 18 21 13 15 19 22 23 24 25 26 26 27 28 29 30 31 32 33 34 35 36 37 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64"
	vals := strings.Split(lst, " ")
	sum := 0
	for _, v := range vals {
		i, _ := strconv.Atoi(v)
		sum += i
	}
	fmt.Println(sum == treeSum(newNodeInt))

	fmt.Println(treeMin(newNodeInt))

	myNode := &NodeInt{
		Val: 5,
		Left: &NodeInt{
			Val: 3,
			Left: &NodeInt{
				Val: 1,
			},
			Right: &NodeInt{
				Val: 4,
			},
		},
		Right: &NodeInt{
			Val: 11,
			Right: &NodeInt{
				Val: 2,
			},
		},
	}

	fmt.Println(maxPathSum(myNode))

}
