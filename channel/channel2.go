package channel

import (
	"fmt"
	"sync"
)

func Channel2() {
	ch := make(chan int, 2)
	ch2 := make(chan int, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		num := <-ch
		fmt.Println("concurrent1: ", num)
	}()

	go func() {
		defer wg.Done()
		num := <-ch
		fmt.Println("concurrent2: ", num)
	}()

	ch <- calc(10)
	ch <- calc(3)

	fmt.Println("ch1 length: ", len(ch))
	fmt.Println("ch2 length: ", len(ch2))

	wg.Wait()
}

func calc(n int) int{
	return n * n
}