package array

import "fmt"

var array = []int{1, 2, 3, 4, 5}

func ArrayInsert() {
	insert(2, 6)
	insert(4, 9)
}

func insert(n int, v int) {
	arr := array
	arr = append(arr[:n+1], arr[n:]...)
	arr[n] = v
	fmt.Println(arr)
}