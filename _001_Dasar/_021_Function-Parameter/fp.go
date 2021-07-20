package main

import "fmt"

// ================================================================================
// Function Parameter:

func sayHello(firstName string, lastName string) { // Parameter
	fmt.Println("Hello", firstName, lastName)
}

// ================================================================================

func main() {
	fmt.Println("===Function Parameter===")

	firstName := "Dliyaulhaq"

	sayHello(firstName, "Mufliansyah")
}
