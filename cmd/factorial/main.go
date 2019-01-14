package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	n, _ := strconv.Atoi(args[1])
	result := factorial(n)
	fmt.Printf("Result: %v\n", result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
