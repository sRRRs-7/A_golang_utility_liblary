package array

import (
	"fmt"
	"testing"
)

type Test struct {
	string
	int
}

func TestMain(t *testing.T) {
	test := send()
	fmt.Println(test.string, test.int)
}

func send() Test {
	return Test{"hello", 1}
}

func TestMain2(t *testing.T) {
	queue := make(chan Test)
	go send1(queue)
	test := <-queue
	fmt.Println(test.string, test.int)
}

func send1(queue chan Test) {
	queue<- Test{"hello", 3}
}

func TestSlice(t *testing.T){
	var arr []string
	arr = []string{"1", "3", "5", "7", "9"}
	arr = append(arr, "hello")
	fmt.Println(arr)
}

func TestInterrupt(t *testing.T){
	arr := []string{"1", "3", "5", "7", "9"}
	arr[0] = "13"
	arr2 := arr[:len(arr)-3]
	arr = append(arr2, arr[len(arr)-2:]...)
	fmt.Println(arr, len(arr))
}

func TestArray(t *testing.T){
	arr := [5]string{"1", "3", "5", "7", "9"}
	fmt.Println(arr)
}

func TestMap(t *testing.T){
	i := 1
	arr := map[int]string{1:"1", 2:"2", i:"3", 4:"4", 5:"5"}
	fmt.Println(arr)
}