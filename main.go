package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"test/all"
)

func main() {
	//channel.ChanMain()
	//linkedList.LinkedList()
	//doubleList.DoubleList()
	//stack.StackMain()
	//stack.LinkedList()
	//queue.QueueMain()
	//queue.LinkedQueue()
	//reverseInput.ReverseInput()
	//parentheses.ParenthesesMain()
	//array.ArrayInsert()
	//tree.LinkTree()
	//tree.NewLinkTree()
	//all.GoAll()
	Trace()
}

func Trace() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	all.GoAll()

	fmt.Println("Hello")
}