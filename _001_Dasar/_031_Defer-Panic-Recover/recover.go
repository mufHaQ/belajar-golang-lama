package main

import "fmt"

func main() {
	fmt.Println("===Recover===")

	runApp(true)
	fmt.Println()
	runApp(false) // Akan menjadi nil
}

func runApp(error bool) {
	fmt.Println("START APP")
	defer endApp()

	if error {
		panic("ERROR")
	}

	// msg := recover() // Salah
	// fmt.Println(msg) // Salah
}

func endApp() {
	msg := recover()

	if msg != nil {
		fmt.Println("Error Message:", msg)
	}

	fmt.Println("END APP")
}
