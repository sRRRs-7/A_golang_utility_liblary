package ioFile

import (
	"bytes"
	"fmt"
	"strings"
)

func MainIo() {
	var b bytes.Buffer
	b.Write([]byte("starbucks"))
	fmt.Println(b.String())

	r := make([]byte, len(b.String()))
	b.Read(r)

	fmt.Println(b.String())
	fmt.Println(string(r))
}

func StringIo() {
	str := "greeeeeeeeeeeeeeeeeen"
	rd := strings.NewReader(str)

	temp := make([]byte, 10)
	cnt, _ := rd.Read(temp)

	fmt.Println(string(temp))
	fmt.Println(string(temp[:cnt]))
}