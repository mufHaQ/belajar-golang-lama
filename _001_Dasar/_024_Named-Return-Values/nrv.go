package main

import "fmt"

func main() {
	fmt.Println("===Named Return Values===")

	fn, _, ln := GetName()
	fmt.Println(fn)
	fmt.Println(ln)
}

func GetName() (firstName, middleName, lastName string) { // Jika ada variable tidak diberi tipe data, maka tipe datanya akan sama dengan variable yang ada dibelakangnya
	firstName = "Dliyaulhaq"
	middleName = ""
	lastName = "Mufliansyah"

	// return firstName, middleName, lastName // Kita bisa langsung mereturn semua data tanpa menyebutkan 1-1 jika sudah ada nama variable dibagian return function
	return
}
