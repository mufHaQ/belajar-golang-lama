package main

import "fmt"

func main() {
	fmt.Println("===If Expression===")

	// ================================================================================
	// If Else Expression:

	// name := "ulhaqs"

	// if name == "ulhaq" {
	// 	fmt.Println("Nama: " + name)
	// } else {
	// 	fmt.Println("Nama Tidak sama")
	// }
	// ================================================================================

	// ================================================================================
	// Else If Expression:

	// var nilaiUjian int8 = 85

	// if nilaiUjian >= 90 {
	// 	fmt.Println("Nilai anda: A")
	// } else if nilaiUjian >= 80 {
	// 	fmt.Println("Nilai anda: B")
	// } else if nilaiUjian >= 70 {
	// 	fmt.Println("Nilai anda: C")
	// } else {
	// 	fmt.Println("Anda GAGAL dalam ujian kali ini, tetap semangat dan coba lagi tahun depan!!")
	// }
	// ================================================================================

	// ================================================================================
	// If dengan Short Statement:

	name := "Dliyaulhaq Mufliansyah"
	// name := "Dliyaulhaq"

	if length := len(name); length > 20 {
		fmt.Println("Maaf, nama anda terlalu panjang")
	} else {
		fmt.Println("Halo " + name)
	}
	// ================================================================================
}
