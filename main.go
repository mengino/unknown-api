package main

import (
	"context"
	"dlsite/app/router"
	"dlsite/internal/config"
	"dlsite/internal/db"
	"dlsite/web/handler/mobile"
	"dlsite/web/handler/pc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-contrib/multitemplate"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var name = "dlsite"

// func main() {
// 	router := gin.Default()

// 	router.Static("/css", "./public/css")
// 	router.Static("/js", "./public/js")
// 	router.Static("/img", "./public/img")
// 	router.Static("/image", "./public/upload")

// 	router.LoadHTMLGlob("public/html/**/*")

// 	router.GET("/index", pc.Index)
// 	router.GET("/mobile/index", mobile.Index)

// 	router.Run(":30088")
// }

func webRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	// 添加两个多模板继承, 初始模板必须写在前面。
	r.AddFromFiles("server", "templates/base.html", "templates/server.html")
	r.AddFromFiles("login", "templates/base.html", "templates/login.html")

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
	// g.Static("/image", "./upload")

	g.Static("/css", "./public/css")
	g.Static("/js", "./public/js")
	g.Static("/img", "./public/img")
	g.Static("/image", "./public/upload")

	// g.LoadHTMLGlob("public/html/**/*")
	g.HTMLRender = webRender()

	g.GET("/index", pc.Index)
	g.GET("/mobile/index", mobile.Index)

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
