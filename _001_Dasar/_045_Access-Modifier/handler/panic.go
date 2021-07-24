package handler

import "fmt"

func PanicHandler() {
	pnc := recover()

	if pnc != nil {
		fmt.Println("Error:", pnc)
	}
}
