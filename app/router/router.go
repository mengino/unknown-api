package router

import (
	"dlsite/app/handler/common"
	"dlsite/app/handler/user"
	"dlsite/app/router/middleware"
	"dlsite/app/router/web"
	"html/template"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func funcMap() template.FuncMap {
	return template.FuncMap{
		"add": func(i int) int {
			return i + 1
		},
	}
}

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
		r.AddFromFilesFuncs(filepath.Base(include), funcMap(), files...)
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

	w := g.Group("")
	{
		web.PCGroup(w)
		web.MobileGroup(w)
	}

	return g
}
