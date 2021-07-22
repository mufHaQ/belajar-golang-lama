package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func (address *Address) GetCityName() string {
	return address.City
}

func (address *Address) ChangeCity(city string) {
	address.City = city
}

func main() {
	fmt.Println("===Pointer di Method===")

	var kudus Address = Address{
		City:     "Kudus",
		Province: "Jawa Tengah",
		Country:  "Indonesia",
	}

	fmt.Println(kudus.GetCityName())

	kudus.ChangeCity("Demak")

	fmt.Println(kudus.GetCityName())
}
