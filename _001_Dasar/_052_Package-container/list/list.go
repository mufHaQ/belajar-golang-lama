package list

import (
	"container/list"
	"fmt"
)

type ListStruct struct {
	List list.List
}

func (ls *ListStruct) CreateList(values ...interface{}) {
	data := list.New()

	for _, value := range values {
		data.PushBack(value)
	}

	ls.List = *data
}

func (ls *ListStruct) GetDataList() list.List {
	return ls.List
}

func (ls *ListStruct) PrintDataList() {
	for dataList := ls.List.Front(); dataList != nil; dataList = dataList.Next() {
		fmt.Println(dataList.Value)
	}
}

func (ls *ListStruct) ListToSlice() []interface{} {
	data := make([]interface{}, 0)

	for dataList := ls.List.Front(); dataList != nil; dataList = dataList.Next() {
		data = append(data, dataList.Value)
	}

	return data
}
