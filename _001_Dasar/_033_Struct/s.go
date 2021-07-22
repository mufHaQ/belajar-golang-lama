package main

import "fmt"

func main() {
	fmt.Println("===Struct===")

	// Membuat Data Struct

	var ulhaq Customer
	ulhaq.Name = "Dliyaulhaq Mufliansyah"
	ulhaq.Adress = "Kudus"
	ulhaq.Age = 16

	fmt.Println(ulhaq)
	fmt.Println(ulhaq.Name)
	fmt.Println(ulhaq.Adress)
	fmt.Println(ulhaq.Age)
}

type Customer struct {
	Name, Adress string
	Age          uint8
}
