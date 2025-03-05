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
	ID          uint
	Name        string
	Price       uint
	Description string
	Soldout     bool
}
