package main

import "fmt"

func main() {
	fmt.Println("===Closure===")

	var counter int
	var name string = "Ulhaq"

	increment := func() string {
		// name = "Budi" // Akan merubah variable name dari scope diatasnya, cara mencegahnya bisa dengan mendeklarasikan ulang
		name := "Budi" // Deklarasi ulang
		fmt.Println("Increment")
		counter++
		return name
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
