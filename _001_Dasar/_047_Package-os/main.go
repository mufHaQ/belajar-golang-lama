package main

import (
	"fmt"
	"os"
	"pos/ospackage"
)

func main() {
	fmt.Println(ospackage.GetArguments())

	host, err := os.Hostname()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(host)
	}

	// fmt.Println(os.Getenv("USERNAME"))
}
