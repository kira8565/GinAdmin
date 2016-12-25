package models

import "github.com/jinzhu/gorm"

type SysRole struct {
	gorm.Model
	RoleName string
	Users    []SysUser  `gorm:"many2many:sys_role_user;"`
}

