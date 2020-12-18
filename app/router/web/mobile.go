package web

import (
	"dlsite/app/handler/web/mobile"

	"github.com/gin-gonic/gin"
)

// MobileGroup 手机端页面路由
func MobileGroup(r *gin.RouterGroup) {
	g := r.Group("mobile")
	{
		g.GET("/", mobile.Index)
		// g.GET("/soft", mobile.Soft)
		// g.GET("/game", mobile.Game)
		// g.GET("/detail/:id", mobile.Detail)
		// g.GET("/news", mobile.News)
		// g.GET("/news/:id", mobile.NewsDetail)
		// g.GET("/hj", mobile.Collection)
		// g.GET("/hj/:id", mobile.CollectionDetail)
		// g.GET("/top", mobile.Top)
	}
}
