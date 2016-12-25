package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	"github.com/kira8565/GinAdmin/models"
)

type IndexController struct {
	db *gorm.DB
}

func NewIndexController(db *gorm.DB) *IndexController {
	return &IndexController{db: db}
}

func (controller IndexController)Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"code":c.DefaultQuery("code", ""),
	})
}

func (controller IndexController)CheckLogin(c *gin.Context) {
	user := &models.SysUser{
		UserName:c.PostForm("username"),
		UserPassword:c.PostForm("password")}
	println(c.PostForm("username"))
	result := controller.db.Where(&user).First(&user)
	if result.RecordNotFound() {
		c.Redirect(http.StatusMovedPermanently, "/?code=1")
	} else {
		println(2)
	}
}
