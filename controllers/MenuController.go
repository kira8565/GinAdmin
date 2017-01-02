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

//菜单管理列表
func (controller MenuController)MenuIndex(c *gin.Context) {
	db := controller.db

	//构造查询条件
	menunameQuery := ""
	if obj, rs := c.GetQuery("menuname"); rs == true {
		menunameQuery = obj
		db = controller.db.Where("menu_name like ?", "%" + menunameQuery + "%")
	}

	//查询分页
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
	db.Find(&list).Count(&total)
	db.Limit(controller.pagesize).Offset(currentPage).Find(&list)

	//返回列表
	res := utils.Paginator(currentPage, controller.pagesize, total)
	c.HTML(http.StatusOK, "menu_index.html", gin.H{
		"menunameQuery":menunameQuery,
		"list":list,
		"paginator":res,
	})
}

