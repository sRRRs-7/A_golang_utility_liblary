package stack

import "fmt"

type Stack []string

func StackMain() {
	s := Stack{}
	s.push("1")
	s.push("2")
	s.push("3")
	fmt.Println(s[0], s[1], s[2])

	for len(s) > 0 {
		el, b := s.pop()
		if b {
			fmt.Println(el)
		}
	}
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() (el string, b bool) {
	if s.isEmpty() {
		return "", false
	}
	index := len(*s)-1
	el = (*s)[index]
	*s = (*s)[:index] // remove the index of the top most element
	return el, true
}