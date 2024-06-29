package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
	num  int
}

type LinkedList struct {
	head *ListNode
}

// 添加节点
func (L *LinkedList) append(value int) {
	newcurr := &ListNode{val: value, next: nil, num: 0}
	if L.head == nil {
		L.head = newcurr
	} else {
		curr := L.head
		for curr.next != nil {
			curr = curr.next
		}
		newcurr.num = curr.num + 1
		newcurr.next = nil
		curr.next = newcurr
	}
}

// 打印节点
func (L *LinkedList) Print() {
	curr := L.head
	for curr != nil {
		fmt.Print(curr.val, " ")
		curr = curr.next
	}
	fmt.Println()
}

// 删除节点
func (L *LinkedList) Delete() {
	if L.head == nil {
		fmt.Println("error,this LinkedList was empty")
		return
	}
	if L.head.next == nil && L.head != nil {
		L.head = nil
		fmt.Print("complete,now this LinkedList was empty")
		return
	}
	curr := L.head
	for curr.next != nil {
		tmp := curr
		curr = curr.next
		if curr.next == nil {
			curr = nil
			tmp.next = nil
			fmt.Println("Delete complete")
			return
		}
	}
}

// 查询节点
func (L LinkedList) find(x int) {
	curr := L.head
	for curr.num != x {
		curr = curr.next
		if curr == nil {
			fmt.Print("find error\n")
			return
		}
	}
	fmt.Printf("The value of num%v is %v\n", x, curr.val)
}

//改节点
func (L *LinkedList) change(x int,y int){
	curr := L.head
	for curr.num != x{
		curr = curr.next
	}
	curr.val = y
	fmt.Println("change completed")
}

func main() {
	var list LinkedList
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.find(2)
	list.find(9)
	list.change(2,89)
	list.Print()
	list.Delete()
	list.Print()
	var list1 LinkedList
	list1.append(1)
	list1.Delete()
	list1.Print()
	var list2 LinkedList
	list2.Delete()
}
