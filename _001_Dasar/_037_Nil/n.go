package main

import "fmt"

func main() {
	fmt.Println("===Nil===")

	var person map[string]string = newMap("Ulhaq")

	if person != nil {
		fmt.Println(person)
	} else {
		fmt.Println("Data Kosong")
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
