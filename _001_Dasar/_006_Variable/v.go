package main

import "fmt"

func main() {
	fmt.Println("===Variable===")

	// ================================================================================
	// Variable:
	// var name string     // Mendeklarasi variable, akan error jika tidak digunakan
	// name = "Dliyaulhaq" // Menambah data ke dalam variable
	// fmt.Println(name)
	// name = "Mufliansyah" // Mengganti isi variable
	// fmt.Println(name)
	// name = 100 // Akan error karena tipe datanya berbeda
	// ================================================================================

	// ================================================================================
	// Tipe Data Variable:
	// var name string = "Dliyaulhaq" // Seperti sebenarnya juga bisa
	// var name = "Dliyaulhaq"
	// fmt.Println(name)
	// var age = 16 // Jika kita tidak menyebutkan tipe datanya, maka akan secara otomatis menggunakan alias 'int', artinya bisa int32/int64
	// fmt.Println(age)
	// ================================================================================

	// ================================================================================
	// Kata Kunci Var:
	// name string := "Dliyaulhaq" // Akan error
	// name := "Dliyaulhaq"
	// fmt.Println(name)
	// name = "Mufliansyah"
	// fmt.Println(name)
	// ================================================================================

	// ================================================================================
	// Deklarasi Multiple Variable:
	var (
		firstName string = "Dliyaulhaq"
		lastName         = "Mufliansyah"
		age       uint8  = 16
	)
	fmt.Println("Hello, my name is", firstName, lastName)
	fmt.Println("Age:", age)
	// ================================================================================
}
