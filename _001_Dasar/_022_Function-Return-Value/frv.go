package main

import "fmt"

func main() {
	fmt.Println("===Function Return Value==")

	name := getName("Dliyaulhaq Mufliansyah")
	fmt.Println(name)

	fmt.Println(getName(""))
}

func getName(name string) string { // Harus mengembalikan string
	// return name

	// Jika menggunakan kondisi yang berbeda-beda
	if name == "" {
		return "UNKNOWN"
	} else {
		return name
	}
}
