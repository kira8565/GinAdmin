package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	"github.com/kira8565/GinAdmin/models"
	"github.com/kira8565/GinAdmin/utils"
)

type MenuController struct {
	db       *gorm.DB
	pagesize int
}

func NewMenuController(db *gorm.DB, pagesize int) *MenuController {
	return &MenuController{db: db, pagesize:pagesize}
}

func (controller MenuController)MenuIndex(c *gin.Context) {
	list := []models.SysMenu{}
	total := int64(0)
	cueerntPage := int(0)
	if page, result := c.Get("page"); result == true {
		cueerntPage = page.(int)
	}
	controller.db.Limit(controller.pagesize).Offset(cueerntPage).Find(&list).Count(&total)
	res := utils.Paginator(cueerntPage, controller.pagesize, total)
	c.HTML(http.StatusOK, "menu_index.html", gin.H{
		"list":list,
		"paginator":res,
	})
}

