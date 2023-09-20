package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Lastname string
	Age      uint64
}

type Session struct {
	gorm.Model
	Id     string
	UserId uint
	User   User
}

type DBVersion struct {
	gorm.Model
	Version uint
}
