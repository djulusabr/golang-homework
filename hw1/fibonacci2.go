package main

import (
	"fmt"
)

func fibonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci2(n - 1) + fibonacci2(n - 2);
}

func main() {
	fmt.Println(fibonacci2(10))
}