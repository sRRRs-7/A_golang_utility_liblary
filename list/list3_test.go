package list

import (
	"fmt"
	"testing"
)

func TestList3(t *testing.T) {
	list := List{}
	list.Insert3(5)
	list.Insert3(4)
	list.Insert3(3)
	list.Insert3(2)
	list.Insert3(1)
	list.Reverse()
	list.Print3()

	var n int
	fmt.Print("Enter a position \n" )
	fmt.Scanf("%d", &n)
	list.Delete(n)
	list.Print3()
}

func (l *List) Insert3(data int) {
	temp := &Node{
		data: data,
		next: l.head,
	}
	l.head = temp
}

func (l *List) Print3() {
	temp := l.head
	for temp != nil {
		fmt.Printf("list is: %d", temp.data)
		temp = temp.next
	}
	fmt.Print()
}

func (l *List) Delete(n int) {
	temp := l.head
	if n == 1 {
		l.head = temp.next // l.head points to second Node
		temp = nil
		return
	}
	for i := 0; i < n-2; i++ {
		temp = temp.next
	}
	// temp points to (n-1)th Node

	temp2 := temp.next // nth Node
	temp.next = temp2.next // (n+1)tn Node
	temp2 = nil
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
