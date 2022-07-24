package tree

import (
	"fmt"
	"math"
)

type Node struct {
	data int
	left *Node
	right *Node
}

// implement binary search tree
func LinkTree() {
	var root *Node = nil

	root = insert(root, 40)
	root = insert(root, 30)
	root = insert(root, 50)
	root = insert(root, 25)
	root = insert(root, 35)
	root = insert(root, 45)
	root = insert(root, 60)

	// Enter Search number
	var number int
	fmt.Print("Enter number be searched: ")
	fmt.Scanf("%d", &number)
	if search(root, number) != nil {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

	// Min, Max
	min := findMin(root)
	fmt.Println("Minimum: ", min.data)

	max := findMax(root)
	fmt.Println("Maximum: ", max.data)

	// Check binary tree
	b := isBinarySearchTree(root, min.data, max.data)
	fmt.Println("Binary search tree ?: ", b)


	// layer count
	h := findHeight(root)
	fmt.Println("Tree height: ", h)

	// Breadth first search
	fmt.Println()
	levelOrder(root)
	fmt.Println()

	// Depth first search
	fmt.Println("PreOrder: ")
	preOrder(root)
	fmt.Println()

	fmt.Println("InOrder: ")
	inOrder(root)
	fmt.Println()

	fmt.Println("PostOrder: ")
	postOrder(root)
	fmt.Println()

	// Delete node
	delete(root, 25)
	levelOrder(root)
	fmt.Println()

	// Successor
	successor := getSuccessor(root, 40)
	fmt.Println(successor)


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

func search(root *Node, data int) *Node {
	if root == nil {return root}
	if root.data == data {return root}
	if data <= root.data {
		return search(root.left, data)
	} else {
		return search(root.right, data)
	}
}

func delete(root *Node, data int) *Node {
	if root == nil {
		fmt.Println("root is nil")
		return root
	} else if data < root.data {
		root.left = delete(root.left, data)
	} else if data > root.data {
		root.right = delete(root.right, data)
	// root.data == data or data not found
	} else {
		// No child
		if root.left == nil && root.right == nil {
			root = nil
		// One child
		} else if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		// Two children
		} else {
			temp := findMin(root.right)
			root.data = temp.data
			root.right = delete(root.right, temp.data)
		}
	}
	return root
}

func findMin(root *Node) *Node {
	if root == nil {
		fmt.Println("Tree is empty")
	}
	curr := root
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func findMax(root *Node) *Node {
	if root == nil {
		fmt.Println("Tree is empty")
	}
	curr := root
	for curr.right != nil {
		curr = curr.right
	}
	return curr
}

// height = edge that from root to leaf
func findHeight(root *Node) float64 {
	if root == nil {
		return -1
	}
	return math.Max(findHeight(root.left), findHeight(root.right)) + 1
}

func levelOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println("level order: ")

	queue := make(chan *Node, 10)
	queue<- root
	for len(queue) != 0 {
		curr := <- queue
		if curr.left != nil {
			queue<- curr.left
		}
		if curr.right != nil {
			queue<- curr.right
		}
		fmt.Println(curr.data)
	}
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	preOrder(root.left)
	preOrder(root.right)
}

// recursion & memory buffer
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root.data)
	inOrder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.data)
}

func isBinarySearchTree(root *Node, min, max int) bool {
	if root == nil {
		return true
	}
	if root.data < min &&
		root.data > max &&
		isBinarySearchTree(root.left, min, root.data) &&
		isBinarySearchTree(root.right, root.data, max) {
			return true
		} else {
			return false
		}
}

func isSubTreeLesser(root *Node, v int) bool {
	if root == nil {
		return true
	}
	if root.data <= v &&
		isSubTreeLesser(root.left, v) &&
		isSubTreeLesser(root.right, v) {
			return true
	} else {
		return false
	}
}

func isSubTreeGreater(root *Node, v int) bool{
	if root == nil {
		return true
	}
	if root.data > v &&
		isSubTreeGreater(root.right, v) &&
		isSubTreeGreater(root.left, v) {
			return true
	} else {
		return false
	}
}

func getSuccessor(root *Node, v int) *Node {
	// search node
	curr := search(root, v)
	if curr == nil {
		return nil
	}

	var successor *Node
	ancestor := root
	// node has right subtree -> find minimum
	if curr.right != nil {
		temp := curr.right
		for temp.left != nil {
			temp = temp.left
			return temp
		}
	// No right subtree
	} else {
		for ancestor != curr {
			if curr.data < ancestor.data {
				successor = ancestor	// deepest node for which current node
				ancestor = ancestor.left
			} else {
				ancestor = ancestor.right
			}
		}
	}
	return successor
}
