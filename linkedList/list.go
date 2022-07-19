package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func LinkedList() {
	list := List{}
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Print()
}

func (l *List) Insert(data int) {
	temp := &Node{
		data: data,
		next: l.head,
	}
	l.head = temp
}

func (l *List) Print() {
	temp := l.head
	fmt.Print("list is: ")
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Print()
}

func (l *List) ReversePrint() {
	temp := l.head
	fmt.Print("list is: ")
	for temp != nil {
		fmt.Printf("%d <- ", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	l.ReversePrint()
}