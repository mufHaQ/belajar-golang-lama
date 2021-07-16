package main

import "fmt"

func main() {
	fmt.Println("===Constant===")

	// ================================================================================
	// Constant:

	// const name string = "Dliyaulhaq"
	// const age = 16 // Tidak error jika tidak digunakan
	// const value = 10000

	// // name = "Test" // Akan error

	// fmt.Println(name)

	// ================================================================================

	// ================================================================================
	// Deklarasi Multiple Constant:

	const (
		firstName string = "Dliyaulhaq"
		lastName         = "Mufliansyah"
		value     int32  = 10000
	)

	fmt.Println("Hello, my name is", firstName, lastName)
	// ================================================================================
}
