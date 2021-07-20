package main

import "fmt"

func GetName(firstName string, middleName string, lastName string) (string, string, string) {
	if middleName == "" {
		return firstName, "", lastName
	} else {
		return firstName, middleName, lastName
	}
}

func main() {
	fmt.Println("===Returning Multiple Values===")

	firstName, _, lastName := GetName("Dliyaulhaq", "", "Mufliansyah") // Kita menggunakan tanda '_' untuk meng-ignore middleName karena hanya mereturn string kosong
	fmt.Println(firstName)
	fmt.Println(lastName)

	// _, _, _ := GetName("", "", "") // Akan terjadi error
}
