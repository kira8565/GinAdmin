package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kira8565/GinAdmin/controllers"
	"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kira8565/GinAdmin/middlewares"
)

func main() {
	r := gin.New()
	store := sessions.NewCookieStore([]byte("GinAdmin"))
	r.Use(sessions.Middleware("GinAdminSession", store))
	r.Use(middlewares.AuthMiddleware())
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/assets", "./assets")

	db, err := gorm.Open("mysql", "root:8565@/ginadmin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	indexController := controllers.NewIndexController(db)
	r.GET("/", indexController.Index)
	r.POST("/checklogin", indexController.CheckLogin)

	mainformController := controllers.NewMainFormController(db)
	r.GET("/mainform", mainformController.MainFormIndex)
	r.GET("/dashboard", mainformController.DashBoardIndex)
	r.Run()
}
