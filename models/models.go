package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Lastname string `gorm:"default:null"`
	Age      uint64
	Test     string
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
