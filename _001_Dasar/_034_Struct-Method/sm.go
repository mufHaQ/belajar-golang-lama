package main

import "fmt"

func main() {
	fmt.Println("===Struct Method===")

	ulhaq := Customer{
		Name:    "Ulhaq",
		Address: "Kudus",
		Age:     16,
	}

	ulhaq.SayHello()
	ulhaq.Test()

	Test(ulhaq)
}

type Customer struct {
	Name, Address string
	Age           int
}

func (cs Customer) SayHello() { // Function ini akan menjadi milik Struct Customer
	fmt.Println("Hello " + cs.Name)
}

func (cs Customer) Test() {
	fmt.Println("Test milik Customer")
}

func Test(cs Customer) { // Tidak akan error walaupun nama functionnya sama
	// fmt.Println("Test")
	fmt.Println(cs)
}
