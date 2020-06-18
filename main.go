package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"golang.org/x/sync/errgroup"
	"lncios.cn/resume/cs"
	"lncios.cn/resume/middleware"
	"lncios.cn/resume/model"
	"lncios.cn/resume/newSession"
	"lncios.cn/resume/rest"
	"lncios.cn/resume/rest/wechat"
	"lncios.cn/resume/router"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	//将数据库拉起
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "lhs", "123321", "39.100.19.104:3306", "resume")
	var err error
	//连接数据库

	cs.Sql, err = xorm.NewEngine("mysql", params)
	cs.Sql.ShowSQL(true)
	if err != nil {
		panic(err)
	}
	//首次运行时加载
	if err := cs.Sql.Sync(
		new(model.Resume)); err != nil {
		fmt.Print("初始化失败", err)

	}

	//启动基础的Http服务
	cs.SessionMgr = newSession.NewSessionMgr("Cookies", 3600)
	app := gin.Default()
	root := app.Group("/api")
	root.Use(middleware.CorsHandler())
	app.Use(middleware.CorsHandler())
	router.Register(root, wechat.User)
	router.Register(root, rest.User)
	router.Register(root, rest.ProjectExprience)
	router.Register(root, rest.WorkExprience)
	router.Register(root, rest.Message)
	router.Register(root, rest.Education)
	server := &http.Server{
		Addr:         "localhost:3001",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})
	fmt.Print("listen:3001")
	if err := g.Wait(); err != nil {
		fmt.Print(err)
	}

}
