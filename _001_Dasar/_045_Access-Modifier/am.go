package main

import (
	"errors"
	"fmt"

	"access-modifier/handler"

	"access-modifier/helper"
)

func main() {
	fmt.Println()

	golang := helper.Application // Bisa diakses

	fmt.Println(helper.SayHello(golang))
	// fmt.Println(helper.sayGoodBye()) // Akan error

	testPanic()
	testError()
}

func testPanic() {
	defer handler.PanicHandler()

	panic("test Panic")
}

func testError() {
	handler.ErrorHandler(errors.New("test Error"))
}
