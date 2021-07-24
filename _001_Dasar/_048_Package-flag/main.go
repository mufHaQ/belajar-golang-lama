package main

import (
	"fmt"
	"pf/db"
)

func main() {
	var db db.DBInterface = &db.Database{}

	db.SetDB()

	for key, val := range db.GetDB() {
		fmt.Println(key, ":", val)
	}
}
