package main

import (
	"fmt"
)

func main() {
	fmt.Println("===Variadic Fucntion===")

	// ================================================================================
	// Variadic Fucntion:

	// fmt.Println(sumAll(10, 10, 10, 10, 10, 5))
	// ================================================================================

	// ================================================================================
	// Slice Parameter:

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumAll(numbers...))
	// ================================================================================
}

func sumAll(numbers ...int) (total int) {
	// total = 0 // kita tidak perlu set nilainya ke 0, karena sudah di set ke int dan default value int adalah 0

	for _, value := range numbers {
		total += value
	}
	return total
}
