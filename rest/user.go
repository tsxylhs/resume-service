package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"lncios.cn/resume/cs"
	"lncios.cn/resume/model"
	"lncios.cn/resume/service"
	"os"
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
		var sessionID = cs.SessionMgr.StartSession(c.Writer, c.Request)

		//设置变量值
		cs.SessionMgr.SetSessionVal(sessionID, "UserInfo", user)
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
func (user) Me(c *gin.Context) {
	id := cs.SessionMgr.GetUserId(c)
	fmt.Print(id)
	role := &model.Role{}
	a := []string{"overview.home", "sys.user", "sys.user_log", "sys.role", "sys.role.edit", "sys.resetPassword"}
	role.Permissions = a
	u := &model.User{}
	u.Roles = append(u.Roles, *role)
	c.JSON(200, u)
}

func (user) download(c *gin.Context) {
	resume := &model.Resume{}
	if err := c.Bind(resume); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
	if resume.Path != "" {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=text")
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(".." + resume.Path)
	} else {
		c.Abort()
		c.JSON(500, "内部服务器错误")
		return
	}
}
func (user) uploadFile(c *gin.Context) {
	resume := model.Resume{}

	form, _ := c.MultipartForm()
	version := c.PostForm("version")
	files := form.File["file[]"]
	resume.Version = version

	// contract.Files=Files
	pathstr := "/resume/" + version
	path := ".." + pathstr
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			c.String(400, "创建文件失败")
			c.Abort()
			return
		} //0777也可以os.ModePerm
		if err := os.Chmod(path, os.ModePerm); err != nil {
			c.String(400, "chmod失败")
			c.Abort()
			return
		}
	}
	fileName := make([]string, 0)
	filePath := make([]string, 0)
	for _, file := range files {
		//文件存储
		str := path + "/" + file.Filename
		strpath := pathstr + "/" + file.Filename
		fWrite, err := os.Create(str)
		if err != nil {
			c.String(500, "服务器错误")
			c.Abort()
			return
		}
		resume.Size = file.Size
		f, _ := file.Open()
		if _, err := io.Copy(fWrite, f); err != nil {

		}
		defer fWrite.Close()
		filePath = append(filePath, strpath)
		fileName = append(fileName, file.Filename)
		resume.Name = file.Filename
		resume.Path = strpath
	}
	c.JSON(200, resume)
}
func (user) getResume(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	resume := &model.Resume{}
	if err := c.Bind(resume); err != nil {
		c.String(400, "参数错误")
		c.Abort()
		return
	}
	r := make(map[string]interface{})
	resumes := &[]model.Resume{}
	if err := service.User.ListResume(page, resume, resumes); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		r["data"] = resumes
		r["page"] = page
		c.JSON(200, r)
	}
}
func (user) saveResume(c *gin.Context) {
	resume := &model.Resume{}
	if err := c.Bind(resume); err != nil {
		c.String(400, "服务器错误")
		c.Abort()
		return
	}
	if err := service.User.SaveResume(resume); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	} else {
		c.JSON(200, resume)
	}
}
func (user) updataResume(c *gin.Context) {
	form := &model.Resume{}
	err := c.Bind(form)

	if err != nil {
		c.String(400, "服务器错误")
		c.Abort()
		return
	}
	if err := service.User.UpdateResume(form); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
	c.String(200, "success")

}
func (user) deleteReusme(c *gin.Context) {
	id := c.Param("id")
	if err := service.User.DeleteResume(id); err != nil {
		c.String(500, "新增失败")
		c.Abort()
		return
	}
	c.String(200, "success")

}
func (user) Register(c *gin.RouterGroup) {
	c.GET("/me", cs.SessionMgr.CheckCookieValid, User.Me)
	c.POST("/v1/web/login", User.login)
	c.POST("/v1/web/create", User.create)
	c.PUT("/v1/web", User.updata)
	c.GET("/v1/web", User.list)
	c.GET("/v1/web/:id", User.get)
	c.DELETE("/v1/web", User.delete)
	c.POST("/v1/resume/upload", User.uploadFile)
	c.POST("/v1/resume/download", User.download)
	c.POST("/v1/resume", User.saveResume)
	c.GET("/v1/resume", User.getResume)
	c.PUT("v1/resume/:id", User.updataResume)
	c.DELETE("/v1/resume/:id", User.deleteReusme)
}
