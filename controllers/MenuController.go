package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/tommy351/gin-sessions"
	"github.com/jinzhu/gorm"
	"github.com/kira8565/GinAdmin/models"
	"github.com/kira8565/GinAdmin/utils"
	"strconv"
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
	currentPage := int(0)
	if page, result := c.GetQuery("page"); result == true {
		pageInt, err := strconv.Atoi(page)
		if (err == nil) {
			currentPage = pageInt - 1
		} else {
			currentPage = 0
		}

	}
	controller.db.Find(&list).Count(&total)
	controller.db.Limit(controller.pagesize).Offset(currentPage).Find(&list)

	res := utils.Paginator(currentPage, controller.pagesize, total)
	c.HTML(http.StatusOK, "menu_index.html", gin.H{
		"list":list,
		"paginator":res,
	})
}

