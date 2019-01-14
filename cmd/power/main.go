package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	a, _ := strconv.Atoi(args[1])
	n, _ := strconv.Atoi(args[2])
	result := power(a, n)
	fmt.Printf("Result: %v\n", result)
}

func power(a, n int) int {
	if n == 0 {
		return 1
	}
	// b <- (a^2)^(n/2)
	b := power(a*a, n/2)
	fmt.Printf("b: %v\n", b)
	if n%2 == 1 {
		b *= a
	}
	return b
}
