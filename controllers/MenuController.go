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

func (controller MenuController)AddMenu(c *gin.Context) {
	c.HTML(http.StatusOK, "menu_add.html", gin.H{
	})
}

func (controller MenuController)AddMenuEntity(c *gin.Context) {
	menu := &models.SysMenu{
		MenuName:c.PostForm("menunane"),
		MenuUrl:c.PostForm("menuurl"),
	}
	controller.db.Save(&menu)
	c.Redirect(http.StatusMovedPermanently, "/menu/index?msg=新增成功&code=alert-success")
}

func (controller MenuController)EditMenu(c *gin.Context) {
	id, rs := c.GetQuery("id")
	if rs == true {
		data := &models.SysMenu{}
		controller.db.First(&data, id)
		c.HTML(http.StatusOK, "menu_edit.html", gin.H{
			"data":data,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/menu/index?msg=系统出错&code=alert-danger")
	}

}

func (controller MenuController)EditMenuEntity(c *gin.Context) {
	id, rs := c.GetPostForm("id")
	if rs == true {
		menu := &models.SysMenu{
			MenuName:c.PostForm("menunane"),
			MenuUrl:c.PostForm("menuurl"),
		}
		controller.db.Model(&models.SysMenu{}).Where("id=?", id).Update(&menu)
		c.Redirect(http.StatusMovedPermanently, "/menu/index?msg=编辑成功&code=alert-success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/menu/index?msg=编辑失败&code=alert-danger")
	}

}

func (controller MenuController)DeleteMenuEntity(c *gin.Context) {
	uid, rs := c.GetQuery("id")
	if rs == true {
		controller.db.Where("id=?", uid).Delete(&models.SysMenu{})
		c.Redirect(http.StatusMovedPermanently, "/menu/index?msg=删除成功&code=alert-success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/menu/index?msg=删除失败,请选择记录&code=alert-danger")
	}

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
	msg := c.DefaultQuery("msg", "")
	code := c.DefaultQuery("code", "")

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
		"msg":msg,
		"code":code,
	})
}

