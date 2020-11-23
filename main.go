package main

import (
	"context"
	"dlsite/app/router"
	"dlsite/internal/config"
	"dlsite/internal/db"
	"dlsite/web/handler/pc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/multitemplate"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var name = "dlsite"

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}

func main() {
	config.Init(name)

	if conn, err := db.Init(); err != nil {
		log.Fatal(err)
	} else {
		defer conn.Close()
		// conn.AutoMigrate(&model.Article{}, &model.Category{}, &model.Product{})
		// return
	}

	g := gin.Default()

	g.Static("/css", "./public/css")
	g.Static("/js", "./public/js")
	g.Static("/img", "./public/img")
	g.Static("/image", "./public/upload")

	g.StaticFile("/favicon.ico", "./public/favicon.ico")

	g.HTMLRender = loadTemplates("public/html/pc")
	// g.HTMLRender = loadTemplates("public/html/mobile")

	g.GET("/", pc.Index)
	g.GET("/soft", pc.Soft)
	g.GET("/game", pc.Game)
	g.GET("/detail", pc.Detail)
	g.GET("/news", pc.News)
	g.GET("/news_detail", pc.NewsDetail)
	g.GET("/hj", pc.Hj)
	g.GET("/hj_detail", pc.HjDetail)
	g.GET("/top", pc.Top)

	// g.GET("/mobile/index", mobile.Index)

	// 启动框架
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(30088),
		Handler:      router.New(g),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
