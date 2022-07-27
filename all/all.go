package all

import (
	"fmt"
)

func GoAll() {
	ch := make(chan int, 3)

	go Multiple(ch, 2)
	go Multiple(ch, 3)
	go Multiple(ch, 4)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func Multiple(ch chan int, num int) {
	result := num * 2
	ch<- result
}