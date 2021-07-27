package main

import (
	"fmt"
	"pc/list"
)

func main() {
	// ================================================================================
	// container/list

	var myList list.ListStruct

	myList.CreateList(1, 2, 3, 4, 5)

	for ls := myList.GetList().Front(); ls != nil; ls = ls.Next() {
		fmt.Println(ls.Value)
	}
	// ================================================================================
}
