package main

import (
	"cmd/fib"
	"fmt"
	"os"
	"strconv"
)

var i int
var fibN int

func main() {
	if len(os.Args) > 1 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			fmt.Printf("%T, %v\n", i, i)
			fibN = fib.Fib(i)
			fmt.Printf("Fibonacci of %d is %d\n", i, fibN)
		}
	}
}
