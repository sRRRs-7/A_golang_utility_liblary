package reverseInput

import (
	"fmt"
	"strings"
)

type Array []string

func ReverseInput() {
	arr := Array{}
	fmt.Print("enter input: ")
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println("input: ", str)
	arr = strings.Split(str, "")
	arr.reverse()
}

func (a *Array) reverse() {
	idx := len(*a) - 1
	for i := 0; i < idx; i++ {
		(*a)[i], (*a)[idx] = (*a)[idx], (*a)[i]
		idx--
	}
	fmt.Println("reverse: ", *a)

	j := strings.Join(*a, "")
	fmt.Println("convert string: ", j)
}