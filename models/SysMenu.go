package models

import "github.com/jinzhu/gorm"

type SysMenu struct {
	gorm.Model
	MenuName string
	Roles    []SysRole  `gorm:"many2many:sys_menu_role;"`
}

