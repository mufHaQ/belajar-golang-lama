package main

import "fmt"

func main() {
	fmt.Println("===Operasi Matematika===")

	// ================================================================================
	// Operasi Matematika:

	// var result = 10 + 10
	// fmt.Println(result)

	// var a = 10
	// var b = 10
	// var c = a * b
	// fmt.Println(c)
	// ================================================================================

	// ================================================================================
	// Augmented Assingments:

	// var i = 10
	// i += 10 // i = i + 10
	// fmt.Println(i)
	// ================================================================================

	// ================================================================================
	// Operator Unary:

	var un int8 = 10
	fmt.Println(un)
	un++ // un = un + 1
	fmt.Println(un)
	un-- // un = un - 1
	fmt.Println(un)

	var pos int8 = 127
	fmt.Println(pos)
	var neg int8 = -128
	fmt.Println(neg)

	var bl = true
	fmt.Println(bl) // true
	bl = !bl
	fmt.Println(bl) // false
	// ================================================================================
}
