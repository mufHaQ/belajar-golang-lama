package main

import "fmt"

func main() {
	fmt.Println("===Operasi Boolean===")

	// ================================================================================
	// Operasi Boolean:

	var (
		ujian        int8 = 90
		absensi      int8 = 74
		lulusUjian   bool = ujian >= 78
		lulusAbsensi bool = absensi >= 75
		lulus        bool = lulusUjian && lulusAbsensi
	)

	fmt.Println(lulus)
	// ================================================================================
}
