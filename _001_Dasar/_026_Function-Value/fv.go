package main

import "fmt"

func main() {
	fmt.Println("===Function Value===")

	type hl func(string) string

	var hello hl = sayHello // Function sebagai value dari variable

	fmt.Println(hello("Dliyaulhaq"))
}

func sayHello(name string) string {
	return "Hello " + name
}
