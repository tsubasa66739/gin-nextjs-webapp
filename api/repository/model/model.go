package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        *uint          `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"default:now()"`
	UpdatedAt time.Time      `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Description string
	Soldout     bool `gorm:"not null"`
}

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
