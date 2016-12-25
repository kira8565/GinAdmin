package models

import "github.com/jinzhu/gorm"

type SysUser struct {
	gorm.Model
	UserName     string
	UserPassword string
}
