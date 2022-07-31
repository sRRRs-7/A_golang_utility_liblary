package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"test/concurrency"
	"test/ioFile"
)

func main() {
	ioFile.Io()
	//channel.ChanMain()
	//channel.Channel2()
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
	//concurrency.GoAll()
	//Trace()
	//stringBuilder.Builder()
	//stringBuilder.Main()
	//design.SyncFunc()
	//design.Main()
	//design.Facets()
	//design.Parameter()
	//design.Function()
	//design.DirectorFunc()
	//factory.FactoryDesign()
	//factory.FactoryDesign2()
	//factory.Factory3()
	//prototype.Prototype()
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

	concurrency.GoAll()

	fmt.Println("Hello")
}