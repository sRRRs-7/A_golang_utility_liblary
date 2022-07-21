package queue

import "fmt"

type Queue []string

func QueueMain() {
	q := Queue{}
	q.push("1")
	q.push("2")
	q.push("3")
	fmt.Println(q[0], q[1], q[2])
	fmt.Println(q[1:])

	for len(q) > 0 {
		el, b := q.pop()
		if b {
			fmt.Println(el)
		}
	}
}

func(q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func(q *Queue) push(str string) {
	*q = append(*q, str)
}

func(q *Queue) pop() (el string, b bool) {
	if q.isEmpty() {
		return "", false
	}
	el = (*q)[0]
	*q = (*q)[1:]
	return el, true
}