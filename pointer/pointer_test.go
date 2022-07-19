package pointer

import (
	"fmt"
	"testing"
)

type TestType struct {
	str string
	n int
}

func TestPointer(t *testing.T) {
	test := TestType{
		str: "test",
		n: 1,
	}
	overwrite(&test)
	fmt.Println(test)
}

func overwrite(test *TestType) {
	test.str = "update test"
	test.n = 2
}