package router

import (
	"dlsite/app/handler/common"
	"dlsite/app/handler/user"
	"dlsite/app/router/middleware"
	"dlsite/web/handler/pc"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

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

// New 初始化路由
func New(g *gin.Engine) *gin.Engine {
	g.Static("/css", "./public/css")
	g.Static("/js", "./public/js")
	g.Static("/img", "./public/img")
	g.Static("/image", "./public/upload")

	g.StaticFile("/favicon.ico", "./public/favicon.ico")

	g.HTMLRender = loadTemplates("public/html/pc")

	g.Use(middleware.Cors())

	r := g.Group("api")
	{
		r.POST("upload", common.Upload)

		u := r.Group("user")
		{
			u.POST("login", user.Login)
		}

		r.Use(middleware.Validate())
		{
			userGroup(r)
			productGroup(r)
			categoryGroup(r)
			articleGroup(r)
		}
	}

	p := g.Group("")
	{
		p.GET("/", pc.Index)
		p.GET("/soft", pc.Soft)
		p.GET("/game", pc.Game)
		p.GET("/detail", pc.Detail)
		p.GET("/news", pc.News)
		p.GET("/news/:id", pc.NewsDetail)
		p.GET("/hj", pc.Hj)
		p.GET("/hj/:id", pc.HjDetail)
		p.GET("/top", pc.Top)
	}

	// m := g.Group("")
	// {
	// 	m.GET("/", pc.Index)
	// 	m.GET("/soft", pc.Soft)
	// 	m.GET("/game", pc.Game)
	// 	m.GET("/detail", pc.Detail)
	// 	m.GET("/news", pc.News)
	// 	m.GET("/news/:id", pc.NewsDetail)
	// 	m.GET("/hj", pc.Hj)
	// 	m.GET("/hj/:id", pc.HjDetail)
	// 	m.GET("/top", pc.Top)
	// }

	return g
}
