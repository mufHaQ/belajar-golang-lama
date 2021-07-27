package main

import (
	"fmt"
	"strconv"
)

func main() {
	parse("1000000000000000000", "int", 10, 64)
	parse("true", "bool")
	parse("100.123", "float", 64)
}

func parse(values ...interface{}) {
	value := values[0].(string)

	valueType := values[1].(string)

	switch valueType {
	case "int":
		baseValue := values[2].(int)
		bitSize := values[3].(int)
		val, _ := strconv.ParseInt(value, baseValue, bitSize)
		fmt.Println(val)
		return
	case "float":
		bitSize := values[2].(int)
		val, _ := strconv.ParseFloat(value, bitSize)
		fmt.Println(val)
		return
	case "bool":
		val, _ := strconv.ParseBool(value)
		fmt.Println(val)
		return
	default:
		fmt.Println("Maaf, terjadi kesalahan!")
		return
	}
}
