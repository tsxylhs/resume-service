package rest

import (
	"github.com/gin-gonic/gin"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"strconv"
)

type message int

var Message message

func (message) create(c *gin.Context) {
	message := &model.Message{}
	if err := c.Bind(message); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.Message.Create(message); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, message)
	}
}
func (message) updata(c *gin.Context) {
	message := &model.Message{}
	if err := c.Bind(message); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.Message.Updata(message); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
}
func (message) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	message := &model.Message{}
	if err := c.Bind(message); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	users := &[]model.Message{}
	if err := service.Message.List(page, message, users); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, users)
	}

}
func (message) delete(c *gin.Context) {
	message := &model.Message{}
	if err := c.Bind(message); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}

	if err := service.Message.Delete(message); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, "删除成功")
	}
}
func (message) get(c *gin.Context) {
	strid := c.Param("id")
	message := &model.Message{}
	message.Id, _ = strconv.ParseInt(strid, 10, 64)
	if err := service.Message.Get(message); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, message)
	}
}
func (message) Register(c *gin.RouterGroup) {
	c.POST("/v1/Message", Message.create)
	c.PUT("/v1/Message", Message.updata)
	c.GET("/v1/Message", Message.list)
	c.GET("/v1/Message/:id", Message.get)
	c.DELETE("v1/Message", Message.delete)
}
