package tree

import (
	"fmt"
)

type Node2 struct {
	data int
	left *Node2
	right *Node2
	parent *Node2
}

// implement binary search tree
func NewLinkTree() {
	var root *Node2 = nil

	root = insert2(root, 20)
	root = insert2(root, 10)
	root = insert2(root, 25)
	root = insert2(root, 7)
	root = insert2(root, 11)
	root = insert2(root, 23)
	root = insert2(root, 28)

	// Enter Search number
	var number int
	fmt.Print("Enter number be searched: ")
	fmt.Scanf("%d", &number)

	node := search2(root, number)
	if node != nil {
		fmt.Println("Found: ", node)
	} else {
		fmt.Println("Not found")
	}

	// print successor
	fmt.Printf("parent of %d is %d ", number, node.parent.data)
}

func newNode2(data int) *Node2 {
	node := &Node2 {
		data: data,
		parent: nil,
		left: nil,
		right: nil,
	}
	return node
}

func insert2(root *Node2, data int) *Node2 {
	if root == nil {
		root = newNode2(data)
	} else if data <= root.data {
		node := insert2(root.left, data)
		root.left = node
		node.parent = root
	} else if data >= root.data {
		node := insert2(root.right, data)
		root.right = node
		node.parent = root
	}
	return root
}

func search2(root *Node2, data int) *Node2 {
	if root == nil {return root}
	if root.data == data {return root}
	if data <= root.data {
		return search2(root.left, data)
	} else {
		return search2(root.right, data)
	}
}

