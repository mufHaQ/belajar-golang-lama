package main

import (
	"fmt"
	"ps/pkg_sort"
	"sort"
)

func main() {
	TestSort()
}

func TestSort() {
	users := pkg_sort.UserSlice{
		{Name: "Budi", Age: 19},
		{Name: "Joko", Age: 21},
		{Name: "Ulhaq", Age: 16},
	}

	fmt.Println(users) // [{Ulhaq 16} {Budi 19} {Joko 21}]

	fmt.Println(users.Less(1, 2)) // False

	sort.Sort(users)

	fmt.Println(users) // [{Ulhaq 16} {Budi 19} {Joko 21}]

	fmt.Println(users.Less(1, 2)) // True

	users.Swap(2, 1)

	fmt.Println(users)
}
