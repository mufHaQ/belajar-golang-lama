package main

import "fmt"

func main() {
	fmt.Println("===Operasi Perbandingan===")

	// ================================================================================
	// Operasi Perbandingan:

	var (
		name1 string = "Ulhaq"
		name2        = "Ulhaq"
		name3        = "Dliya"
		name4        = "Ulhaq"

		result1 bool = (name1 == name2)
		result2      = name3 != name4
		result3      = (len(name3) > len(name4))
	)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	// ================================================================================
}
