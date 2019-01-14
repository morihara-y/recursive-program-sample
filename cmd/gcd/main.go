package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	a, _ := strconv.Atoi(args[1])
	b, _ := strconv.Atoi(args[2])
	var result int
	if a > b {
		result = gcd(a, b)
	} else {
		result = gcd(b, a)
	}
	fmt.Printf("Result: %v\n", result)
}

func gcd(a, b int) int {
	if b == 0 {
		result := math.Abs(float64(a))
		return int(result)
	}
	return gcd(b, a%b)
}
