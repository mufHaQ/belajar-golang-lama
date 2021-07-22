package main

import "fmt"

func main() {
	fmt.Println("===Function as Parameter===")

	sayHelloWithFilter("Anjing", filterName)
	sayHelloWithFilter("Ulhaq", filterName)
}

type Filter (func(string) string)

// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter Filter) { // Menggunakan Type Declarations
	fmt.Println("Hello " + filter(name))
}

func filterName(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
