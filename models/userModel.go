package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Company   string
	Email     string `gorm:"unique"`
	Password  string
}
