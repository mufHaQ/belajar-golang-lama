package main

import "fmt"

func main() {
	fmt.Println("===Konversi Tipe Data===")

	// ================================================================================
	// Konversi Tipe Data(1):

	// var (
	// 	nilai32 int32 = 2147483647
	// 	nilai64 int64 = int64(nilai32)

	// 	nilai16 int16 = int16(nilai32) // Tidak akan muat karena ukurannya terlalu besar, dan terjadi integer overflow
	// )

	// fmt.Println(nilai32) // 2147483647
	// fmt.Println(nilai64) // 2147483647
	// fmt.Println(nilai16) // -1
	// ================================================================================

	// ================================================================================
	// Konversi Tipe Data(2):

	var (
		name    string = "Dliyaulhaq"
		e       uint8  = name[0]
		eString string = string(e)
	)

	fmt.Println(name)
	fmt.Println(eString)
	// ================================================================================
}
