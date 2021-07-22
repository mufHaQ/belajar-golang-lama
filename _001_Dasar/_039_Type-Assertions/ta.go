package main

import (
	"fmt"
)

func main() {
	defer panicHandler()

	fmt.Println("===Type Assertions===")

	random("String")
	random(100)
	random(true)
}

func random(data interface{}) {
	switch value := data.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown Type")
	}
}

func panicHandler() {
	msg := recover()

	if msg != nil {
		fmt.Println("Error:", msg)
	} else {
		return
	}
}
