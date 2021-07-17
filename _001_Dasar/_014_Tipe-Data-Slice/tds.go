package main

import (
	"fmt"
)

func main() {
	fmt.Println("===Tipe Data Slice===")

	// ================================================================================
	// Tipe Data Slice:

	// var months = [...]string{ // Memakai '...' jika kita tidak tahu berapa kapasitas Arraynya
	// 	"Januari",
	// 	"Februari",
	// 	"Maret",
	// 	"April",
	// 	"Mei",
	// 	"Juni",
	// 	"Juli",
	// 	"Agustus",
	// 	"September",
	// 	"Oktober",
	// 	"November",
	// 	"Desember",
	// }

	// var slice1 = months[4:7]
	// ================================================================================

	// ================================================================================
	// Function Pada Slice (len, cap):

	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// slice1[0] = "Diubah"
	// fmt.Println(months) // 'Mei' di Array 'months' akan diubah
	// fmt.Println(slice1)
	// ================================================================================

	// ================================================================================
	// Function Pada Slice (append):

	// days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// daySlice1 := days[3:6]
	// daySlice1[0] = "'Kamis Baru'"
	// // fmt.Println(daySlice1)
	// fmt.Println("days:", days)

	// daySlice2 := append(daySlice1, "'Minggu Baru'") // Data 'Minggu' pada Array 'days' akan diganti
	// fmt.Println("daySlice2:", daySlice2)
	// fmt.Println("days:", days)

	// daySlice3 := days[3:]                         // Kita menggunakan semua kapasitas
	// daySlice4 := append(daySlice3, "'Hari Baru'") // Karena kapasitas penuh, maka akan dibuat Array baru
	// fmt.Println("daySlice4:", daySlice4)
	// daySlice4[0] = "aksnDLKSNDLKANSD" // Tidak ada data yang diganti pada Array 'days'
	// fmt.Println("days:", days)        // Data Array 'days' masih sama dengan Slice 'daySlice2'
	// ================================================================================

	// ================================================================================
	// Function Pada Slice (make):

	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "Dliyaulhaq"
	// newSlice[1] = "Mufliansyah"

	// newSlice = append(newSlice, "Budi")
	// newSlice = append(newSlice, "Test")

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))
	// ================================================================================

	// ================================================================================
	// Function Pada Slice (copy):

	// fromSlice := make([]int8, 3, 100)
	// fromSlice[0] = 1
	// fromSlice[1] = 2
	// fromSlice[2] = 3

	// toSlice := make([]int8, len(fromSlice), cap(fromSlice)) // Jika length lebih kecil dari source slice, maka akan otomatis terpotong

	// copy(toSlice, fromSlice)

	// fmt.Println(fromSlice)
	// fmt.Println(toSlice)
	// ================================================================================

	// ================================================================================
	// Hati-Hati Saat Membuat Array:

	iniArray := [...]string{"...", "..."} // Atau juga '...' diganti angka
	iniSlice := []string{"...", "..."}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	// ================================================================================
}
