package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string `gorm:"primaryKey"`
	UserTickets uint
}
