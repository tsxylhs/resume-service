package rest

import (
	"github.com/gin-gonic/gin"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"strconv"
)

type projectExprience int

var ProjectExprience projectExprience

func (projectExprience) create(c *gin.Context) {
	projectExprience := &model.ProjectExprience{}
	if err := c.Bind(projectExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.ProjectExprience.Create(projectExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, projectExprience)
	}
}
func (projectExprience) updata(c *gin.Context) {
	projectExprience := &model.ProjectExprience{}
	if err := c.Bind(projectExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.ProjectExprience.Updata(projectExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
}
func (projectExprience) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	projectExprience := &model.ProjectExprience{}
	if err := c.Bind(projectExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	users := &[]model.ProjectExprience{}
	if err := service.ProjectExprience.List(page, projectExprience, users); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, users)
	}

}
func (projectExprience) delete(c *gin.Context) {
	projectExprience := &model.ProjectExprience{}
	if err := c.Bind(projectExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}

	if err := service.ProjectExprience.Delete(projectExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, "删除成功")
	}
}
func (projectExprience) get(c *gin.Context) {
	strid := c.Param("id")
	projectExprience := &model.ProjectExprience{}
	projectExprience.Id, _ = strconv.ParseInt(strid, 10, 64)
	if err := service.ProjectExprience.Get(projectExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, projectExprience)
	}
}
func (projectExprience) Register(c *gin.RouterGroup) {
	c.POST("/v1/project", ProjectExprience.create)
	c.PUT("/v1/project", ProjectExprience.updata)
	c.GET("/v1/project", ProjectExprience.list)
	c.GET("/v1/project/:id", ProjectExprience.get)
	c.DELETE("v1/project", ProjectExprience.delete)
}
