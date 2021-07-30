package main

import (
	"fmt"
	"pc/list"
	"pc/ring"
)

func Wrapper(str string, fn func()) {
	fmt.Println("========================================")
	fmt.Println(str + ":")
	fn()
	fmt.Println("========================================")
}

func main() {
	// ================================================================================
	// container/list
	Wrapper("container/list", ContainerList)
	// ================================================================================

	// ================================================================================
	// container/ring
	Wrapper("container/ring", ContainerRing)
	// ================================================================================
}

func ContainerList() {
	var myList list.ListStruct

	myList.CreateList(1, 2, 3, 4, 5)

	// myList.PrintDataList()

	myListToSlice := myList.ListToSlice()

	fmt.Println("List to Slice:", myListToSlice)
}

func ContainerRing() {
	ring.TestRing(5, "ring ke-")
}
