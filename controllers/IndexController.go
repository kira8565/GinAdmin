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
	//s := sessions.Get(c)
	//s.Set("a", "d")
	//s.Save()
	//println(s.Get("a").(string))
	user := models.User{}
	rs := controller.db.Find(&user)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"rs":rs,
	})
}
