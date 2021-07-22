package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func ChangeCity(address *Address, kota string) {
	address.Kota = kota
}

func main() {
	fmt.Println("===Pointer di Function===")

	address := Address{"Kudus", "Jawa Tengah", "Indonesia"}

	fmt.Println(address) // {Kudus Jawa Tengah Indonesia}

	ChangeCity(&address, "Demak")

	fmt.Println(address) // {Demak Jawa Tengah Indonesia}
}
