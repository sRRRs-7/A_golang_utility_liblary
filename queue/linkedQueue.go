package queue

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
	prev *Node
}

type List struct {
	front *Node
	rear *Node
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

	if l.front != nil {
		l.front.prev = node
	}
	l.front = node

	f := l.front
	for f.next != nil {
		f = f.next
	}
	l.rear = f
}

func (l *List) deQueue() (node *Node, b bool) {
	if l.isEmpty() {
		return node, false
	}
	node = l.rear
	l.rear.prev.next = nil
	l.rear = nil
	l.rear = node.prev
	return node, true
}

func (l *List) isEmpty() bool {
	return l.front == nil
}