package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	n, _ := strconv.Atoi(args[1])
	table := map[int]int{}
	result := fibonacci(n, table)
	fmt.Printf("Result: %v\n", result)
}

func fibonacci(n int, table map[int]int) int {
	value, contain := table[n]
	if contain {
		return value
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	result := fibonacci(n-1, table) + fibonacci(n-2, table)
	table[n] = result
	return result
}
