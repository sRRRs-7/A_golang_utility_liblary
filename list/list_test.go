package list

import (
	"fmt"
	"testing"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func TestList(t *testing.T) {
	var list = List{}
	fmt.Println("How many nodes?")
	var n, i, x int
	fmt.Scanf("%d", &n)
	for i = 0; i < n; i++ {
		fmt.Printf("Enter the number \n")
		fmt.Scanf("%d", &x)
		list.Insert(x)
		list.Print()
	}
}

func (l *List) Insert(x int) {
	temp := &Node{
		data: x,
		next: l.head,
	}
	l.head = temp
}

func (l *List) Print() {
	var temp *Node
	temp = l.head
	fmt.Printf("List is: ")
	for temp != nil {
		fmt.Printf("%d", temp.data)
		temp = temp.next
	}
	fmt.Printf("\n")
}
