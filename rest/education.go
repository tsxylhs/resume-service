package rest

import (
	"github.com/gin-gonic/gin"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"strconv"
)

type education int

var Education education

func (education) create(c *gin.Context) {
	education := &model.Education{}
	if err := c.Bind(education); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.Education.Create(education); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, education)
	}
}
func (education) updata(c *gin.Context) {
	education := &model.Education{}
	if err := c.Bind(education); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.Education.Updata(education); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
}
func (education) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	education := &model.Education{}
	if err := c.Bind(education); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	users := &[]model.Education{}
	if err := service.Education.List(page, education, users); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, users)
	}

}
func (education) delete(c *gin.Context) {
	education := &model.Education{}
	if err := c.Bind(education); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}

	if err := service.Education.Delete(education); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, "删除成功")
	}
}
func (education) get(c *gin.Context) {
	strid := c.Param("id")
	education := &model.Education{}
	education.Id, _ = strconv.ParseInt(strid, 10, 64)
	if err := service.Education.Get(education); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, education)
	}
}
func (education) Register(c *gin.RouterGroup) {
	c.POST("/v1/education", Education.create)
	c.PUT("/v1/education", Education.updata)
	c.GET("/v1/education", Education.list)
	c.GET("/v1/education/:id", Education.get)
	c.DELETE("v1/education", Education.delete)
}
