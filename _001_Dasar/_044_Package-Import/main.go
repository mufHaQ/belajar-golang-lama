package main

import (
	"belajar-import/helper"
	"belajar-import/others"
)

func main() {
	defer helper.PanicHandler()

	others.SayHello()

	panic("Test")
}
