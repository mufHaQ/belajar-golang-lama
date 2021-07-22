package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("===Error Interface===")

	// var contohError error = errors.New("Ups Error")
	// fmt.Println(contohError)

	hasil, err := Pembagian(100, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasil)
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan 0")
	} else {
		return nilai / pembagi, nil
	}
}
