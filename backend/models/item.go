package models

import "gorm.io/gorm"

type Item struct {
    gorm.Model
    UserID      uint   // Foreign key ke tabel users
    Name        string `gorm:"not null"`
    Description string
}
