package channel

import (
	"fmt"
)

func ChanMain() {
	n := 3.0
	n2 := 2

	ch := make(chan float64)
	ch2 := make(chan int)

	// stack data structure therefore process2 finished before process1
	go process1(n, ch)
	go process2(n2, ch2)

	select {
	case res := <-ch:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	}
}

// generics
func process1[T1 float64 | int](n T1, ch chan<- T1) {
	result1 := n * n
	ch<- result1
}

func process2[T1 float64 | int](n T1, ch2 chan<- T1) {
	result1 := n * n
	ch2<- result1
}