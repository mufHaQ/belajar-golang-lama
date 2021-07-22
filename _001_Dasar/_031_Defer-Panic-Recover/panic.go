package main

import "fmt"

func main() {
	fmt.Println("===Panic===")
	// runPanic(false)
	runPanic(true)
}

func runPanic(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("NO ERROR")
}

func endApp() {
	fmt.Println("END APP")
}
