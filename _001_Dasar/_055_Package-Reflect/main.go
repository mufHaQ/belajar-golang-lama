package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{
		Name: "Ulhaq",
	}

	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(sampleType)
	fmt.Println(structField)
}
