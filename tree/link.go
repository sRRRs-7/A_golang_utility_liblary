package tree

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func LinkTree() {
	var root *Node = nil

	root = insert(root, 15)
	root = insert(root, 12)
	root = insert(root, 16)
	root = insert(root, 14)
	root = insert(root, 23)
	root = insert(root, 1)
	root = insert(root, 8)

	fmt.Println(root)
	fmt.Println(root.left)
	fmt.Println(root.right)

	var number int
	fmt.Print("Enter number be searched: ")
	fmt.Scanf("%d", &number)
	if search(root, number) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}

func newNode(data int) *Node {
	node := &Node {
		data: data,
		left: nil,
		right: nil,
	}
	return node
}

func insert(root *Node, data int) *Node {
	if root == nil {
		root = newNode(data)
	} else if data <= root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}
	return root
}

func search(root *Node, data int) bool {
	if root == nil {return false}
	if root.data == data {return true}
	if data <= root.data {
		return search(root.left, data)
	} else {
		return search(root.right, data)
	}
}