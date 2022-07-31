package design

import "fmt"

type instance struct {
	Count int
}

var count int

func SyncFunc() {
	i := instance{}
	i.onceInstance()

	// addition & subtraction 2
	count = 2
	add2()
	add2()
	subtract2()
	fmt.Println(count)
}

func (i *instance) onceInstance() {
	sum1 := i.add()
	sum2 := i.add()
	fmt.Println(sum1, sum2)

	subtract1 := i.subtract()
	subtract2 := i.subtract()
	fmt.Println(subtract1, subtract2)
}
func (i *instance) add() int {
	i.Count += 1
	return i.Count
}
func (i *instance) subtract() int {
	i.Count -= 1
	return i.Count
}


func add2() {
	count += 1
}
func subtract2() {
	count -= 1
}