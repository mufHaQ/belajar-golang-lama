package main

import "fmt"

func main() {
	fmt.Println("===Tipe Data Map===")

	// ================================================================================
	// Tipe Data Map:

	// person := map[string]string{
	// 	"name": "Dliyaulhaq Mufliansyah",
	// 	"city": "Kudus",
	// }

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["city"])
	// ================================================================================

	// ================================================================================
	// Function Pada Map (len, map[key]):

	// names := map[string]string{
	// 	"firstName": "Dliyaulhaq",
	// 	"lastName":  "Mufliansyah",
	// }

	// fmt.Println(len(names))
	// fmt.Println(names["firstName"])
	// ================================================================================

	// ================================================================================
	// Function Pada Map (make, delete):

	person := make(map[string]string)

	person["firstName"] = "Dliyaulhaq"
	person["lastName"] = "Mufliansyah"
	person["ups"] = "Salah"

	fmt.Println(person)

	delete(person, "ups") // person["ups"] akan dihapus

	fmt.Println(person)
	// ================================================================================
}
