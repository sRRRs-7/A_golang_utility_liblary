package ioFile

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Scan() {
	str, err := getInput(os.Stdin, flag.Args()...)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(str)
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	fmt.Print("Enter: ")
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("empty input is not allowed")
	}
	return text, nil
}

