package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func InsertAtEnd(head, node *Node) *Node {
	if head == nil {
		return node
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node

	return head
}

func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func FromSlice(vals []int) *Node {
	var head *Node
	for _, v := range vals {
		head = InsertAtEnd(head, &Node{Val: v})
	}

	return head
}

//	n				n
//
// c				c
// 1 -> 2 -> 3, 1 -> 2 -> 3, 1 -> 2 -> 3
// p				p
func ReverseIterative(head *Node) *Node {
	var prev *Node

	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func ReverseRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func MiddleNode(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
