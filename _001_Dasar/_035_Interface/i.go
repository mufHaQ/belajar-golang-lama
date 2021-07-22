package main

import "fmt"

func main() {
	fmt.Println("===Interface===")

	var ulhaq Makhluk = Person{
		Name: "Ulhaq",
	}

	fmt.Println(ulhaq)
}

type Makhluk interface {
	GetName() string
}

type Person struct {
	Name string
}

func (p Person) GetName() string { // Otomatik mengimplementasikan Interface Makhluk
	return p.Name
}
