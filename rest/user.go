package rest

import (
	"github.com/gin-gonic/gin"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"strconv"
)

type user int

var User user

func (user) login(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if user, ok := service.User.Login(user); ok {
		c.JSON(200, user)
	} else {
		c.String(500, "用户信息错误")
		c.Abort()
		return
	}

}
func (user) create(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.User.Create(user); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, user)
	}
}
func (user) updata(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	if err := service.User.Updata(user); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
}
func (user) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	users := &[]model.User{}
	if err := service.User.List(page, user, users); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, users)
	}

}
func (user) delete(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}

	if err := service.User.Delete(user); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, "删除成功")
	}
}
func (user) get(c *gin.Context) {
	strid := c.Param("id")
	user := &model.User{}
	user.Id, _ = strconv.ParseInt(strid, 10, 64)
	if err := service.User.Get(user); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, user)
	}
}
func (user) Register(c *gin.RouterGroup) {
	c.POST("/v1/web/login", User.login)
	c.POST("/v1/web/create", User.create)
	c.PUT("/v1/web", User.updata)
	c.GET("/v1/web", User.list)
	c.GET("/v1/web/:id", User.get)
	c.DELETE("v1/web", User.delete)
}
