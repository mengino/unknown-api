package main

import (
	"context"
	"dlsite/app/router"
	"dlsite/internal/config"
	"dlsite/internal/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var name = "dlsite"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("public/html/**/*")

	router.GET("/pc/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pc/index.html", "pc")
	})
	router.GET("/mobile/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mobile/index.html", "mobile")
	})

	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.Static("/img", "./public/img")
	router.Static("/image", "./public/upload")

	router.Run(":30088")

	// r := gin.Default()
	// r.LoadHTMLGlob("html/**/*")
	// r.GET("/index", func(c *gin.Context) {
	// 	c.HTML(200, "pc/index.html", "flysnow_org")
	// })
	// r.GET("/m", func(c *gin.Context) {
	// 	c.HTML(200, "mobile/index.html", "aaa")
	// })
	// r.Run(":30088")
}

func main1() {
	config.Init(name)

	if conn, err := db.Init(); err != nil {
		log.Fatal(err)
	} else {
		defer conn.Close()
		// conn.AutoMigrate(&model.Article{}, &model.Category{}, &model.Product{})
		// return
	}

	g := gin.Default()
	g.Static("/image", "./upload")

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
