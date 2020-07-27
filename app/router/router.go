package router

import (
	"dlsite/app/handler/common"
	"dlsite/app/handler/user"
	"dlsite/app/router/middleware"

	"github.com/gin-gonic/gin"
)

// New 初始化路由
func New(g *gin.Engine) *gin.Engine {
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

	return g
}
