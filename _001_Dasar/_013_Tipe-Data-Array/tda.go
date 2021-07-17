package main

import "fmt"

func main() {
	fmt.Println("===Tipe Data Array===")

	// ================================================================================
	// Tipe Data Array:

	// var names [2]string      // Di dalam [] adalah panjang Arraynya
	// names[0] = "Dliyaulhaq"  // Index 0
	// names[1] = "Mufliansyah" // Index 1
	// // names[2] = "Test"        // Akan terjadi error karena melebihi kapasitas, yang seharusnya 2 diisi 3

	// fmt.Println(names)
	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// ================================================================================

	// ================================================================================
	// Membuat Array Langsung Dengan Data:

	// var values = [3]int {80, 90, 100} // Ini bisa
	// var values = [3]int{ // Seperti ini juga bisa
	// 	80,
	// 	90,
	// 	100, // Tambah koma di bagian akhir
	// }
	// fmt.Println(values)
	// ================================================================================

	// ================================================================================
	// Function Pada Array:

	var names = [2]string{
		"Dliyaulhaq",
		"Mufliansyah",
	}
	fmt.Println(len(names))
	fmt.Println(names[1])

	names[0] = "Berubah"

	fmt.Println(names[0])
	fmt.Println(names)
	// ================================================================================
}
