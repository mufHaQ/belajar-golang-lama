package main

import (
	"fmt"
	"time"
)

func main() {
	// Now()

	// Manual()

	Parsing()
}

func Now() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
}

func Manual() {
	local := time.Date(2020, 10, 14, 18, 0, 0, 0, time.Local)
	fmt.Println(local)
}

func Parsing() {
	parse, _ := time.Parse(time.RFC1123, "Mon, 02 Jan 2006 15:04:05 MST")
	fmt.Println(parse)
}
