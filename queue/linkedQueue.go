package queue

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type List struct {
	front *Node
}

func LinkedQueue() {
	l := List{}
	l.enQueue("1")
	l.enQueue("2")
	l.enQueue("3")
	l.enQueue("4")

	node, b := l.deQueue()
	if b {
		fmt.Println("deQueue: ", node.data)
	}

	fmt.Print("all: ")
	for l.front != nil {
		fmt.Print(l.front.data)
		l.front = l.front.next
	}
}

func (l *List) enQueue(data string) {
	node := &Node {
		data: data,
		next: l.front,
	}
	l.front = node
}

func (l *List) deQueue() (node *Node, b bool) {
	if l.isEmpty() {
		return node, false
	}
	node = l.front
	l.front = l.front.next
	return node, true
}

func (l *List) isEmpty() bool {
	return l.front == nil
}