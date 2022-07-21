package stack

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
	prev *Node
}

type List struct {
	head * Node
	tail *Node
}

func LinkedList() {
	list := List{}
	list.push("1")
	list.push("2")
	list.push("3")
	list.push("4")

	temp := list.head
	for temp != nil {
		fmt.Printf("%s ", temp.data)
		temp = temp.next
	}
	fmt.Println()

	node, b := list.stackPop()
	if b {
		fmt.Println("stackPop: ", node.data)
	}
	node, b = list.queuePop()
	if b {
		fmt.Println("queuePop: ", node.data)
	}

	list.print()
}

func (l *List) print() {
	temp := l.head
	for temp != nil {
		fmt.Println("rest value: ", temp.data)
		temp = temp.next
	}
}

func (l *List) isEmpty() bool {
	return l.head == nil || l.tail == nil
}

func (l *List) push(data string) {
	temp := &Node{
		data: data,
		next: l.head,
	}
	if l.head != nil {
		l.head.prev = temp
	}
	l.head = temp

	h := l.head
	for h.next != nil {
		h = h.next
	}
	l.tail = h
}

func (l *List) stackPop() (node *Node, b bool) {
	if l.isEmpty(){
		return nil, false
	}
	node = l.tail
	l.tail.prev.next = nil
	l.tail = nil
	l.tail = node.prev
	return node, true
}

func (l *List) queuePop() (node *Node, b bool) {
	if l.isEmpty(){
		return nil, false
	}
	node = l.head
	l.head.next.prev = nil
	l.head = nil
	l.head = node.next
	return node, true
}

