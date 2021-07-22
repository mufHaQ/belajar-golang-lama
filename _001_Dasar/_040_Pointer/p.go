package main

import "fmt"

func main() {
	fmt.Println("===Pointer===")

	// ================================================================================
	// Pass by Value:

	fmt.Println("\nPass-by-Value")
	passByValue()
	// ================================================================================

	// ================================================================================
	// Pass by Reference | Operator &:

	fmt.Println("\nPass-by-Reference")
	passByReference()
	// ================================================================================

	// ================================================================================
	// Operator *:

	fmt.Println("\nOperator Tanpa Bintang (*)")
	OperatorTanpaBintang()
	fmt.Println("\nOperator Dengan Bintang (*)")
	OperatorDenganBintang()
	// ================================================================================

	// ================================================================================
	// Function new:

	fmt.Println("\nFunction new")
	functionNew()
	// ================================================================================
}

type Address struct {
	Kota, Provinsi, Negara string
}

func passByValue() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := address1

	address2.Kota = "Demak"

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
}

func passByReference() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2.Kota = "Demak"

	fmt.Println(address1) // address1 akan berubah
	fmt.Println(address2)
}

func OperatorTanpaBintang() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2.Kota = "Demak"

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // {Demak Jawa Tengah Indonesia} | address1 tidak berubah
	fmt.Println(address2) // &{Jakarta DKI Jakarta Indonesia}
}

func OperatorDenganBintang() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	var address3 *Address = &address1

	address2.Kota = "Demak"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1) // {Malang Jawa Timur Indonesia} | data address1 akan berubah menjadi data dari address2
	fmt.Println(address2) // &{Malang Jawa Timur Indonesia}
	fmt.Println(address3) // &{Malang Jawa Timur Indonesia}
}

func functionNew() {
	var address1 *Address = new(Address)
	var address2 *Address = address1

	// Kita harus set data 1-1
	address2.Kota = "Kudus"
	address2.Provinsi = "Jawa Tengah"
	address2.Negara = "Indonesia"

	fmt.Println(address1) // address1 berubah
	fmt.Println(address2)
}
