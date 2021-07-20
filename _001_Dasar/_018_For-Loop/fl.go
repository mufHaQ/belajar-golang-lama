package main

import "fmt"

func main() {
	fmt.Println("===For Loops===")

	// ================================================================================
	// For Loops:

	// var counter int8 = 1

	// for counter <= 100 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }
	// Kode program seperti ini biasanya terdapa di while loop jika dibahasa pemrogramman lain
	// ================================================================================

	// ================================================================================
	// For Dengan Statement:

	//    Init(1)   |   Kondisi(2)  |    Post(4)
	// for counter := 1; counter <= 100; counter++ {
	// 	fmt.Println("Perulangan Ke", counter) // Program (3)
	// }
	// ================================================================================

	// ================================================================================
	// Mengakses Data Collection Menggunakan Perulangan:

	// names := []string{
	// 	"Dliyaulhaq",
	// 	"Mufliansyah",
	// 	"Budi",
	// 	"Joko",
	// }

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	// ================================================================================

	// ================================================================================
	// For Range:

	// Array
	fmt.Println("Array")
	names := [...]string{
		"Dliyaulhaq",
		"Mufliansyah",
	}

	for index, value := range names {
		fmt.Println("Index", index, ":", value)
	}

	fmt.Println()

	// Slice
	fmt.Println("Slice")
	person := names[:]
	for _, value := range person { // karena range mereturn 2 nilai, kita bisa menggunakan tanda '_' jika data tidak digunakan
		fmt.Println("Nama:", value)
	}

	fmt.Println()

	// Map
	fmt.Println("Map")
	address := map[string]string{
		"street":  "Kudus-Purwodadi",
		"city":    "Kudus",
		"country": "Indonesia",
	}

	for key, val := range address {
		fmt.Println("Key", key, ":", val)
	}
	// ================================================================================
}
