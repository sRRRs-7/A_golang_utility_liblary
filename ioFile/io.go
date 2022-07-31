package ioFile

import (
	"fmt"
	"log"
	"os"
)

func Io() {
	// write
	f2, err := os.Create("text.txt")
	if err != nil {
		log.Fatalf("cannot create file: %v", err)
	}
	defer func() {
		err := f2.Close()
		if err != nil {
			log.Fatalf("cannot close file: %v", err)
		}
	}()
	str := "world"
	data := []byte(str)
	cnt, err := f2.Write(data)
	if err != nil {
		log.Fatalf("cannot write file: %v", err)
	}
	fmt.Printf("%d bytes\n", cnt)

	// process file id
	fd := f2.Fd()
	fmt.Println(fd)


	// read
	f, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(string(f))
}