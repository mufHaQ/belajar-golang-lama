package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "!Dliyaulhaq Mufliansyah!"

	fmt.Println(strings.Contains(name, "ulhaq")) // true
	fmt.Println(strings.Split(name, " "))        // [!Dliyaulhaq, Mufliansyah!]
	fmt.Println(strings.ToLower(name))           // !dliyaulhaq mufliansyah!
	fmt.Println(strings.ToUpper(name))           // !DLIYAULHAQ MUFLIANSYAH!
	fmt.Println(strings.Trim(name, "!"))         // DLIYAULHAQ MUFLIANSYAH
	fmt.Println(strings.ReplaceAll(name, strings.Trim(name, "!"), "Mufliansyah Dliyaulhaq"))

	fmt.Println(strings.Trim("a    "+name+"    a", " ")) // Spasi tidak akan hilang karena trim hanya menghilangkan bagian kiri dan kanan saja, karena ada huruf 'a' jadi tidak dihilangkan
}
