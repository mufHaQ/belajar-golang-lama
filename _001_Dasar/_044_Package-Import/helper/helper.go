package helper

import "fmt"

func PanicHandler() {
	msg := recover()

	if msg != nil {
		fmt.Println("Error:", msg)
	}
}
