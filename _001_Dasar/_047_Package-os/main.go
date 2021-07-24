package main

import (
	"fmt"
	"pos/ospackage"
)

func main() {
	fmt.Println(ospackage.GetArguments())
	fmt.Println(ospackage.GetHostname())

	// fmt.Println(os.Getenv("USERNAME"))
}
