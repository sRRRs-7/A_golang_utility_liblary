package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("text.txt")
	if err != nil {
		log.Fatalf("could not create text.txt: %v", err)
	}
	defer f.Close()

	n, err := f.Write([]byte("hello world"))
	if err != nil {
		log.Fatalf("could not write text.txt: %v", err)
	}

	fmt.Println("text.txt is written: ", n, "bytes")

	// standard output
	os.Stdout.Write([]byte("hello world"))
	fmt.Println()

	// buffer
	var b bytes.Buffer
	b.Write([]byte("buffer storage"))
	fmt.Println("buffer: ", b.String())
}