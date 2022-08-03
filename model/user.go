package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID       string `gorm:"unique;not null" json:"uuid"`
	Email      string `gorm:"uniqueIndex;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	IsVerified bool   `gorm:"default:false" json:"is_verified"`
}

type Phone struct {
	gorm.Model
	UUID       string `gorm:"unique" json:"uuid"`
	Number     string `gorm:"" json:"number"`
	IsVerified bool   `gorm:"default:false" json:"is_verified"`
}
