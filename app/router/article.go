package router

import (
	"dlsite/app/handler/article"

	"github.com/gin-gonic/gin"
)

func articleGroup(r *gin.RouterGroup) {
	g := r.Group("article")
	{
		g.GET("", article.List)
		g.GET("/:id", article.Detail)
		g.POST("", article.Create)
		g.PUT("/:id", article.Update)
		g.DELETE("", article.Delete)
	}
}
