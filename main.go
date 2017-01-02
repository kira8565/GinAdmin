package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kira8565/GinAdmin/middlewares"
	"github.com/kira8565/GinAdmin/commons"
)

func main() {

	r := gin.New()
	r.Static("/assets", "./assets")
	store := sessions.NewCookieStore([]byte("GinAdmin"))
	r.Use(sessions.Middleware("GinAdminSession", store))
	r.Use(middlewares.AuthMiddleware())
	r.LoadHTMLGlob("templates/**/*")

	db, err := gorm.Open("mysql", "root:8565@/ginadmin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	commons.BindRouters(r, db)
	r.Run()
}
