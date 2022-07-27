package main

import (
	"fmt"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	fmt.Println(b.N)
}