package list

import (
	"container/list"
	"fmt"
)

type ListStruct struct {
	data *list.List
}

func (ls *ListStruct) CreateList(values ...interface{}) {
	data := list.New()

	for _, value := range values {
		data.PushBack(value)
	}

	ls.data = data
}

func (ls *ListStruct) PrintList() {
	for lst := ls.data.Front(); lst != nil; lst = lst.Next() {
		fmt.Println(lst.Value)
	}
}

func (ls *ListStruct) GetList() *list.List {
	return ls.data
}
