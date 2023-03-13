package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"unique"`
	Password       string 
	Role           string 
	ProfilePicture string
	Phone          string
	Address        string
}

