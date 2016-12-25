package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/kira8565/GinAdmin/models"
)

type MainFormController struct {
	db *gorm.DB
}

func NewMainFormController(db *gorm.DB) *MainFormController {
	return &MainFormController{db: db}
}

func (controller MainFormController)MainFormIndex(c *gin.Context) {
	menus := []models.SysMenu{}
	controller.db.Find(&menus)

	c.HTML(http.StatusOK, "mainform.tmpl", gin.H{
		"menus":menus,
	})
}

func (controller MainFormController)DashBoardIndex(c *gin.Context) {
	menus := []models.SysMenu{}
	controller.db.Find(&menus)

	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
	})
}