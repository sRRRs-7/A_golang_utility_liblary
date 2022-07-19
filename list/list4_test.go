package list

import (
	"fmt"
	"testing"
)

func TestList4(t *testing.T) {
	list := List{}
	list.Insert4(5)
	list.Insert4(4)
	list.Insert4(3)
	list.PrintRecursion()
}

func (l *List) Insert4(data int) {
	temp := &Node{
		data: data,
		next: l.head,
	}
	l.head = temp
}

func (l *List) PrintRecursion() {
	head := l.head
	if head == nil {
		return
	}
	fmt.Printf("%d", head.data)
	fmt.Print(head.next) // recursion call
}
