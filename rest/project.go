package rest

import (
	"github.com/gin-gonic/gin"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"strconv"
	"strings"
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
	r := make(map[string]interface{})
	users := &[]model.ProjectExprience{}
	if err := service.ProjectExprience.List(page, projectExprience, users); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		r["data"] = users
		r["page"] = page
		c.JSON(200, r)
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
func (projectExprience) uploadImage(c *gin.Context) {
	file, h, err := c.Request.FormFile("file")
	if err != nil {

		c.String(400, "未找到file文件")
		c.Abort()
		return
	}
	form := &model.File{}
	form.OriginName = h.Filename[:strings.Index(h.Filename, ".")]
	form.Suffix = h.Filename[strings.Index(h.Filename, "."):]
	// 保存成为新的记录
	if err := service.ProjectExprience.UploadFile(form, file); err != nil {

		c.String(500, "保存失败")
		c.Abort()
		return
	}
	// 保存成功后返回新的记录
	c.JSON(200, form)
}
func (projectExprience) Register(c *gin.RouterGroup) {
	c.POST("/v1/project", ProjectExprience.create)
	c.PUT("/v1/project/:id", ProjectExprience.updata)
	c.GET("/v1/project", ProjectExprience.list)
	c.GET("/v1/project/:id", ProjectExprience.get)
	c.DELETE("v1/project", ProjectExprience.delete)
	c.POST("/v1/file", ProjectExprience.uploadImage)
}
