package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	fmt.Println("===Struct===")

	// ================================================================================
	// Membuat Data Struct:

	// var ulhaq Customer
	// ulhaq.Name = "Dliyaulhaq Mufliansyah"
	// ulhaq.Address = "Kudus"
	// ulhaq.Age = 16

	// fmt.Println(ulhaq)
	// fmt.Println(ulhaq.Name)
	// fmt.Println(ulhaq.Address)
	// fmt.Println(ulhaq.Age)
	// ================================================================================

	// ================================================================================
	// Struct Literals:

	ulhaq := Customer{ // Ini bisa
		Name:    "Dliyaulhaq Mufliansyah",
		Address: "Kudus",
		Age:     16,
	}

	ulhaq2 := Customer{"Ulhaq", "Kudus", 16} // Ini bisa

	fmt.Println(ulhaq)
	fmt.Println(ulhaq2)
	// ================================================================================
}
