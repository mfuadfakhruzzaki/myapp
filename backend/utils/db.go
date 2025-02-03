package utils

import (
	"fmt"
	"log"

	"github.com/mfuadfakhruzzaki/myapp-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    // Ganti string koneksi sesuai kebutuhan Anda
    dsn := "host=localhost user=postgres password=020803 dbname=myappdb port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error membuka database dengan GORM: ", err)
    }

    fmt.Println("Berhasil terhubung ke database dengan GORM!")

    // Auto migration: membuat atau mengubah tabel berdasarkan model
    err = DB.AutoMigrate(&models.User{}, &models.Item{})
    if err != nil {
        log.Fatal("Error dalam auto migration: ", err)
    }
    fmt.Println("Auto migration selesai, tabel siap digunakan!")
}
