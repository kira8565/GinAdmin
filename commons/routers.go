package commons

import (
	"github.com/gin-gonic/gin"
	"github.com/kira8565/GinAdmin/controllers"
	"github.com/jinzhu/gorm"
)

func BindRouters(r *gin.Engine, db *gorm.DB) {
	indexController := controllers.NewIndexController(db)
	r.GET("/", indexController.Index)
	r.POST("/checklogin", indexController.CheckLogin)

	mainformController := controllers.NewMainFormController(db)
	r.GET("/mainform", mainformController.MainFormIndex)
	r.GET("/dashboard", mainformController.DashBoardIndex)

	menuController := controllers.NewMenuController(db)
	r.GET("/menu/index", menuController.MenuIndex)
}
