package rest

import (
	"github.com/gin-gonic/gin"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"strconv"
)

type workExprience int

var WorkExprience workExprience

func (workExprience) create(c *gin.Context) {
	workExprience := &model.WorkExprience{}
	if err := c.Bind(workExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.WorkExprience.Create(workExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, workExprience)
	}
}
func (workExprience) updata(c *gin.Context) {
	workExprience := &model.WorkExprience{}
	if err := c.Bind(workExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.WorkExprience.Updata(workExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
}
func (workExprience) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	workExprience := &model.WorkExprience{}
	if err := c.Bind(workExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	r := make(map[string]interface{})
	users := &[]model.WorkExprience{}
	if err := service.WorkExprience.List(page, workExprience, users); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		r["data"] = users
		r["page"] = page
		c.JSON(200, r)
	}
}
func (workExprience) delete(c *gin.Context) {
	workExprience := &model.WorkExprience{}
	if err := c.Bind(workExprience); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}

	if err := service.WorkExprience.Delete(workExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, "删除成功")
	}
}
func (workExprience) get(c *gin.Context) {
	strid := c.Param("id")
	workExprience := &model.WorkExprience{}
	workExprience.Id, _ = strconv.ParseInt(strid, 10, 64)
	if err := service.WorkExprience.Get(workExprience); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, workExprience)
	}
}
func (workExprience) Register(c *gin.RouterGroup) {
	c.POST("/v1/work", WorkExprience.create)
	c.PUT("/v1/work/:id", WorkExprience.updata)
	c.GET("/v1/work", WorkExprience.list)
	c.GET("/v1/work/:id", WorkExprience.get)
	c.DELETE("v1/work", WorkExprience.delete)
}
