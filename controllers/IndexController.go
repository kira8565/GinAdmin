package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	"github.com/kira8565/GinAdmin/models"
	"github.com/tommy351/gin-sessions"
)

type IndexController struct {
	db *gorm.DB
}

func NewIndexController(db *gorm.DB) *IndexController {
	return &IndexController{db: db}
}

func (controller IndexController)Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"msg":c.DefaultQuery("msg", ""),
	})
}

func (controller IndexController)CheckLogin(c *gin.Context) {

	user := &models.SysUser{
		UserName:c.PostForm("username"),
		UserPassword:c.PostForm("password")}
	println(c.PostForm("username"))
	result := controller.db.Where(&user).First(&user)
	if result.RecordNotFound() {
		c.Redirect(http.StatusMovedPermanently, "/?msg=用户名密码有误")
	} else {
		session := sessions.Get(c)
		session.Set("isLogin", "true")
		session.Set("userName", user.UserName)
		session.Set("userId", user.ID)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/mainform")
	}
}
