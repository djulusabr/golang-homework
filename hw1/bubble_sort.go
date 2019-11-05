package main

import (
	"fmt"
)

func bubble_sort(a *[]float64) {
	n := len((*a))
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if (*a)[j] > (*a)[j + 1] {
				(*a)[j], (*a)[j + 1] = (*a)[j + 1], (*a)[j]
			}
		}
	}
}

func main() {
	a := []float64{8, 7, 6, 5, 3, 4, 1, 2}
	
	printSlice(a)
	bubble_sort(&a)
	printSlice(a)
}

func printSlice(s []float64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}