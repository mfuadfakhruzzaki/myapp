package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    connStr := "host=localhost port=5432 user=postgres password=020803 dbname=myappdb sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error membuka database: ", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("Tidak dapat terhubung ke database: ", err)
    }
    fmt.Println("Berhasil terhubung ke database!")
}
