package main

import "fmt"

func main() {
	fmt.Println("===Defer===")
	// deferRunApp()
	deferRunAppWithError(0) // Function logging akan tetap dijalankan walaupun error
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func deferRunApp() {
	defer logging() // Akan dijalankan terakhir
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println("Run Application")

}

func deferRunAppWithError(value int) {
	defer logging() // Akan dijalankan terakhir

	fmt.Println(10 / value)

	// fmt.Println("Run Application") // Tidak akan dieksekusi jika terjadi error pada kode diatasnya
}
