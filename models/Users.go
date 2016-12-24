package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName     string
	UserPassword string
}
