package database

import "fmt"

var (
	Connection string
)

func init() {
	Connection = "MYSQL" // Ketika file ini diakses, function ini akan otomatis dijalankan
	fmt.Println(Connection)
}

func GetDatabase() string {
	return Connection
}
