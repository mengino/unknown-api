package router

import (
	"dlsite/app/handler/category"

	"github.com/gin-gonic/gin"
)

func categoryGroup(r *gin.RouterGroup) {
	g := r.Group("category")
	{
		g.GET("", category.List)
		g.GET("/:id", category.Detail)
		g.POST("", category.Create)
		g.PUT("/:id", category.Update)
		g.DELETE("", category.Delete)
	}
}
