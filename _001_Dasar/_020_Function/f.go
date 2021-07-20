package main

import "fmt"

func main() {
	fmt.Println("===Function===")

	// ================================================================================
	// Function:

	for index := 1; index <= 100; index++ {
		sayHello()
	}
	// ================================================================================
}

func sayHello() {
	fmt.Println("Hello World")
}
