package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"golang.org/x/sync/errgroup"
	"lncios.cn/resume/cs"
	"lncios.cn/resume/middleware"
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
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "lib", "lib123", "47.103.212.217:3306", "libDemo")
	var err error
	//连接数据库
	cs.Sql, err = xorm.NewEngine("mysql", params)
	if err != nil {
		panic(err)
	}
	//首次运行时加载
	//model.NewBD()
	//启动基础的Http服务
	cs.SessionMgr = newSession.NewSessionMgr("Cookies", 3600)
	app := gin.Default()
	root := app.Group("/api")
	root.Use(middleware.CorsHandler())
	router.Register(root, wechat.User)
	router.Register(root, rest.User)

	server := &http.Server{
		Addr:         ":3001",
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
