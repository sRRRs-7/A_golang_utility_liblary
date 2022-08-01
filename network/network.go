package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		Server()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		defer wg.Done()
		Client()
	}()

	wg.Wait()
}

// write = send
// read = receive
func Server() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	conn, err := lis.Accept()
	if err != nil {
		log.Fatalf("Could not accept: %v", err)
	}

	str := "I lone MacBook"
	data := []byte(str)
	cnt, err := conn.Write(data)
	if err != nil {
		log.Fatalf("Could not write: %v", err)
	}
	fmt.Println(cnt, "bytes")
}

func Client() {
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Could not dial connection: %v", err)
	}

	data := make([]byte, 1024)
	cnt, err := conn.Read(data)
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}
	fmt.Println(string(data[:cnt]))
}