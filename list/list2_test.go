package list

import (
	"fmt"
	"testing"
)

func TestList2(t *testing.T) {
	list := List{}
	list.Insert2(2, 1)
	list.Insert2(3, 2)
	list.Insert2(4, 1)
	list.Insert2(5, 2)
	list.Print2()
}

// keyword: node, buffer
func (l *List) Insert2(data, n int) {
	temp := &Node{
		data: data,
		next: nil,
	}
	if n == 1 {
		temp.next = l.head
		l.head = temp
		return
	}

	var prevTemp = l.head
	for i := 0; i < n-2; i++ {
		prevTemp = prevTemp.next
	}
	temp.next = prevTemp.next
	prevTemp.next = temp
}

func (l *List) Print2() {
	var temp *Node
	temp = l.head
	fmt.Printf("List is: ")
	for temp != nil {
		fmt.Printf("%d", temp.data)
		temp = temp.next
	}
	fmt.Printf("\n")
}