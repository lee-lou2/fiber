package model

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	UserID      int
	User        User `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhoneID     int
	Phone       Phone  `gorm:"references:ID" json:"phone"`
	UUID        string `gorm:"unique;not null" json:"uuid"`
	Number      string `gorm:"uniqueIndex;not null" json:"number"`
	Description string `gorm:"" json:"description"`
	IsDefault   bool   `gorm:"default:false" json:"is_default"`
}
