package main

import "fmt"

func main() {
	fmt.Println("===Recursive Function===")

	fmt.Println(factorialRecursive(5))
	fmt.Println(factorialLoop(4))
}

func factorialLoop(val int) (result int) {
	result++
	for ; val > 0; val-- {
		result *= val
	}
	return
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
