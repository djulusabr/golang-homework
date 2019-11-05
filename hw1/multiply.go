package main

import (
	"fmt"
)

func multiply(a, b float64) float64 {
	if a == 1 {
		return b
	}
	if a == 0 {
		return 0
	}
	if a < 0 {
		return -multiply(-a, b)
	}
	return b + multiply(a - 1, b)
}

func main() {
	fmt.Println(multiply(1, -5))
}