package main

import "fmt"

func Wrapper(str string, fn func()) {
	fmt.Println("========================================")
	fmt.Println(str + ":")
	fn()
	fmt.Println("========================================", "\n")
}

func main() {
	// ================================================================================
	// Pass by Value
	Wrapper("Pass by Value", PassByValue)
	// ================================================================================

	// ================================================================================
	// Pass by Reference
	Wrapper("Pass by Reference", PassByReference)
	// ================================================================================

	// ================================================================================
	// Without Asterisk
	Wrapper("Without Asterisk", WithoutAsterisk)
	// ================================================================================

	// ================================================================================
	// With Asterisk
	Wrapper("With Asterisk", WithAsterisk)
	// ================================================================================

	// ================================================================================
	// Function New
	Wrapper("Function New", New)
	// ================================================================================
}

type Address struct {
	City, Province, Country string
}

func PassByValue() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := address1 // Data dari address1 diduplikasi dan disimpan di address2

	address2.City = "Demak"

	fmt.Println("address1:", address1) // {Kudus Jawa Tengah Indonesia} | Harusnya berubah jika dibahasa seperti PHP/JavaScript
	fmt.Println("address2:", address2) // {Demak Jawa Tengah Indonesia}
}

func PassByReference() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1 // Operator Ampersand (&)

	address2.City = "Demak"

	fmt.Println("address1:", address1) // {Demak Jawa Tengah Indonesia}
	fmt.Println("address2:", address2) // {Demak Jawa Tengah Indonesia}
}

func WithoutAsterisk() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2.City = "Pati" // Hanya mengubah fieldnya

	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"} // address2 akan membuat alamat pointer yang baru dan bukan ke address1

	fmt.Println("address1:", address1) // {Pati Jawa Tengah Indonesia}
	fmt.Println("address2:", address2) // &{Malang Jawa Timur Indonesia}
}

func WithAsterisk() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Pati"

	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"} // Operator Asterisk | address1 akan berpindah ke alamat yang sama milik address2

	fmt.Println("address1:", address1)
	fmt.Println("address2:", address2)
	fmt.Println("address3:", address3)
}

func New() {
	address1 := new(Address)
	address1.City = "Kudus"
	address1.Province = "Jawa Tengah"
	address1.Country = "Indonesia"

	address2 := address1

	fmt.Println("address1:", address1)
	fmt.Println("address2:", address2)
}
