package main

import "fmt"

func main() {
	fmt.Println("===Interface Kosong===")

	dataString := interfaceKosong("Data String")
	fmt.Println(dataString)

	dataInt := interfaceKosong(10000)
	fmt.Println(dataInt)

	dataBool := interfaceKosong(true)
	fmt.Println(dataBool)
}

type Any interface{}

func interfaceKosong(data Any) interface{} {
	return data
}
