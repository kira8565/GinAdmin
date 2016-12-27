package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	"github.com/kira8565/GinAdmin/models"
)

type MenuController struct {
	db *gorm.DB
}

func NewMenuController(db *gorm.DB) *MenuController {
	return &MenuController{db: db}
}

func (controller MenuController)MenuIndex(c *gin.Context) {
	list := []models.SysMenu{}
	total := 0
	controller.db.Find(&list).Count(&total)
	c.HTML(http.StatusOK, "menu_index.html", gin.H{
		"list":list,
		"total":total,
	})
}
