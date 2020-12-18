package web

import (
	"dlsite/app/handler/web/pc"

	"github.com/gin-gonic/gin"
)

// PCGroup PC端页面路由
func PCGroup(r *gin.RouterGroup) {
	r.GET("/", pc.Index)
	r.GET("/soft", pc.Soft)
	r.GET("/game", pc.Game)
	r.GET("/detail/:id", pc.Detail)
	r.GET("/news", pc.News)
	r.GET("/news/:id", pc.NewsDetail)
	r.GET("/hj", pc.Collection)
	r.GET("/hj/:id", pc.CollectionDetail)
	r.GET("/top", pc.Top)
}
