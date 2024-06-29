package main

import (
	"fmt"
	"time"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (root *Node) Insert(val int) {
	if root == nil {
		return
	}

	if val < root.Value {
		if root.Left == nil {
			root.Left = &Node{Value: val}
		} else {
			root.Left.Insert(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &Node{Value: val}
		} else {
			root.Right.Insert(val)
		}
	}
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	//fmt.Println(root.Value)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func main() {
	time_start := time.Now()

	root := &Node{Value: 5}

	for i := 1; i < 2000; i++ {
		root.Insert(i)
	}
	//PreOrderTraversal(root)

	time_end := time.Now()
	all_time := time_end.Sub(time_start)
	fmt.Print(all_time)
}
