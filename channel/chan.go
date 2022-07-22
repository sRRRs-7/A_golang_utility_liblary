package channel

import (
	"fmt"
	"math"
)

type Number interface {
	~float64 | ~int
}

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
func process1[T1 Number](n T1, ch chan<- float64) {
	result := math.Pow(float64(n), float64(2))
	ch<- result
}

func process2[T1 float64 | int](n T1, ch2 chan<- T1) {
	result := n * n
	ch2<- result
}