package main

import (
	"fmt"
	"math"
)

func main() {
	var val float64 = 100.5
	var val2 float64 = 200.5

	fmt.Println(math.Round(val))
	fmt.Println(math.Floor(val))
	fmt.Println(math.Ceil(val))
	fmt.Println(math.Max(val, val2))
	fmt.Println(math.Min(val, val2))
}
