package main

import "fmt"

func main() {
	fmt.Println("===Switch Expression===")

	// ================================================================================
	// Switch Expression:

	// name := "Ulhaq"

	// switch name {
	// case "Dliya":
	// 	fmt.Println("Halo " + name)
	// case "Ulhaq":
	// 	fmt.Println("Halo " + name)
	// default:
	// 	fmt.Println("Halo, boleh kenalan?")
	// }
	// ================================================================================

	// ================================================================================
	// Switch Expression dengan Short Statement:

	// var name string = "Ulhaq"
	// // name = "Dliyaulhaq Mufliansyah"

	// switch length := len(name); length > 20 {
	// case true:
	// 	fmt.Println("Nama anda terlalu panjang!")
	// case false:
	// 	fmt.Println("Nama anda sudah benar")
	// default:
	// 	fmt.Println("Maaf, terjadi kesalahan")
	// }
	// ================================================================================

	// ================================================================================
	// Switch Tanpa Kondisi:

	name := [...]string{
		"Dliyaulhaq",
		"Mufliansyah",
		"sadasd",
		"sadasd",
		"asdasd",
	}

	length := len(name)

	switch {
	case length >= 5:
		fmt.Println("Maaf, nama anda terlalu panjang")
	case length < 5:
		fmt.Println("Selamat datang!")
	default:
		fmt.Println("Maaf, terjadi kesalahan")
	}
	// ================================================================================
}
