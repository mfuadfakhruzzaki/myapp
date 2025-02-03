package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    var err error
    DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
    if err != nil {
        log.Fatal("Error membuka database dengan GORM: ", err)
    }
    fmt.Println("Berhasil terhubung ke database!")
}
