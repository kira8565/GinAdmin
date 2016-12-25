package main

import (
	"github.com/jinzhu/gorm"
	"github.com/kira8565/GinAdmin/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:8565@/ginadmin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DropTable(
		&models.SysUser{},
		&models.SysRole{},
		&models.SysMenu{},
		"ginadmin.sys_role_user",
		"ginadmin.sys_menu_role")

	db.AutoMigrate(
		&models.SysUser{},
		&models.SysRole{},
		&models.SysMenu{})

	adminUser := &models.SysUser{UserName:"1", UserPassword:"1"}
	db.Create(adminUser)

	adminRole := &models.SysRole{RoleName:"管理员", Users:[]models.SysUser{*adminUser}}
	db.Create(adminRole)

	indexMenus := &models.SysMenu{
		MenuName:"首页",
	}
	db.Create(&indexMenus)

	menuMenu := &models.SysMenu{
		MenuName:"菜单管理", ParentId:indexMenus.ID,
	}

	db.Create(&menuMenu)
}
