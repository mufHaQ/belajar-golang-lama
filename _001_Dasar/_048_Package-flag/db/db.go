package db

import "flag"

type DBInterface interface {
	GetDB() map[string]string
	SetDB()
}

type Database struct {
	DBHOST, DBNAME, DB, DBUSER, DBPASS string
}

var (
	DBHOST, DBNAME, DB, DBUSER, DBPASS *string
)

func init() {
	DBHOST = flag.String("host", "localhost", "Masukkan host database, defaultnya adalah localhost")    // Nilai defaultnya adalah localhost
	DBNAME = flag.String("dbname", "mysql", "Masukkan nama database yang anda pakai, defaultnya mysql") // Nilai defaultnya adalah mysql
	DB = flag.String("db", "", "Masukkan table mana yang akan anda pakai")
	DBUSER = flag.String("user", "", "Masukkan username database")
	DBPASS = flag.String("pass", "", "Masukkan password database")

	flag.Parse() // Kita perlu memanggil function flag.Parse(), jika tidak maka data yang kita masukkan tidak akan dimasukkan dan akan menggunakan nilai default
}

func (db *Database) SetDB() {
	db.DBHOST = *DBHOST
	db.DBNAME = *DBNAME
	db.DBUSER = *DBUSER
	db.DBPASS = *DBPASS
	db.DB = *DB
}

func (db *Database) GetDB() map[string]string {
	return map[string]string{
		"DBHOST": db.DBHOST,
		"DBNAME": db.DBNAME,
		"DB    ": db.DB,
		"DBUSER": db.DBUSER,
		"DBPASS": db.DBPASS,
	}
}
