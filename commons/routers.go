package commons

import (
	"github.com/gin-gonic/gin"
	"github.com/kira8565/GinAdmin/controllers"
	"github.com/jinzhu/gorm"
)

func BindRouters(r *gin.Engine, db *gorm.DB) {
	GlobalPageSize := 10
	indexController := controllers.NewIndexController(db)
	r.GET("/", indexController.Index)
	r.POST("/checklogin", indexController.CheckLogin)

	mainformController := controllers.NewMainFormController(db)
	r.GET("/mainform", mainformController.MainFormIndex)
	r.GET("/dashboard", mainformController.DashBoardIndex)

	menuController := controllers.NewMenuController(db, GlobalPageSize)
	r.GET("/menu/index", menuController.MenuIndex)
	r.GET("/menu/add", menuController.AddMenu)
	r.GET("/menu/delete", menuController.DeleteMenuEntity)
	r.POST("/menu/addeneity", menuController.AddMenuEntity)
}
