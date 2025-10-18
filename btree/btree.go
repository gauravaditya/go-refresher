package btree

import (
	"fmt"
	"practise-go/datatype/stack"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val}
}

func (n *Node) appendLeft(l *Node) *Node {
	n.Left = l

	return n
}

func (n *Node) appendRight(r *Node) *Node {
	n.Right = r

	return n
}

// InOrder traverses a tree in "Left, Root, Right" fashion.
func InOrder(node *Node) {
	if node == nil {
		return
	}

	InOrder(node.Left)
	fmt.Println(node.Value)
	InOrder(node.Right)
}

// PreOrder traverses a tree in "Root, Left, Right" fashion.
func PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// PostOrder traverses a tree in "Left, Right, Root" fashion.
func PostOrder(node *Node) {
	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
}

// InOrderIterative performs an in-order traversal without recursion.
func InOrderIterative(root *Node) {
	// Your implementation here
	stakk := stack.New[*Node]()
	current := root

	// Loop continues as long as there's a node to process OR nodes on the stack.
	for current != nil || !stakk.IsEmpty() {

		// Phase 1: Find the leftmost node of the current subtree.
		for current != nil {
			stakk.Push(current) // Push current node to stack
			current = current.Left
		}

		// At this point, current is nil.
		// It's time to visit the node at the top of the stack.
		// This is the leftmost node we haven't visited yet.

		// Phase 2: Pop, Visit, Go Right.
		popped, ok := stakk.Pop()
		if !ok {
			continue
		}

		fmt.Println(popped.Value)
		current = popped.Right
	}
}

//	        8
//	     /     \
//	    4       12
//	   / \     /  \
//	  2   6   10   14
//	 / \ / \ / \  / \
//	1  3 5  7 9 11 13 15
//
// 8 -> 4 -> 2 -> 1 -> 3 -> 6 -> 5 -> 7 -> 12 -> 10 -> 9 -> 11 -> 14 -> 13 -> 15
// PreOrderIterative performs a pre-order traversal without recursion.
func PreOrderIterative(root *Node) {
	stakk := stack.New[*Node]()
	current := root

	for current != nil || !stakk.IsEmpty() {

		if current == nil {
			var ok bool
			current, ok = stakk.Pop()
			if !ok {
				break
			}
		}
		fmt.Println(current.Value)

		if current.Right != nil {
			stakk.Push(current.Right)
		}

		current = current.Left
	}
}

// PostOrderIterative performs a post-order traversal without recursion.
func PostOrderIterative(root *Node) {
	stakk := stack.New[*Node]()
	current := root

	var lastVisited *Node
	for current != nil || !stakk.IsEmpty() {
		for current != nil {
			stakk.Push(current)
			current = current.Left
		}

		peeked, ok := stakk.Peek()
		if !ok {
			break
		}

		if peeked.Right != nil && peeked.Right != lastVisited {
			current = peeked.Right
			continue
		}

		popped, ok := stakk.Pop()
		if !ok {
			break
		}

		fmt.Println(popped.Value)
		lastVisited = popped
	}
}

// LevelOrder performs a breadth-first (level-by-level) traversal.
func LevelOrder(root *Node) {
	q := []*Node{root}

	for len(q) > 0 {
		node := q[0]
		fmt.Println(node.Value)

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		q = q[1:]
	}
}
