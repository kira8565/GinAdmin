package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
)

type MenuController struct {
	db *gorm.DB
}

func NewMenuController(db *gorm.DB) *MenuController {
	return &MenuController{db: db}
}

func (controller MenuController)MenuIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "menu_index.tmpl", gin.H{
	})
}
