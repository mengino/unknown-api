package router

import (
	"dlsite/app/handler/product"

	"github.com/gin-gonic/gin"
)

func productGroup(r *gin.RouterGroup) {
	g := r.Group("product")
	{
		g.GET("", product.List)
		g.GET("/:id", product.Detail)
		g.POST("", product.Create)
		g.PUT("/:id", product.Update)
		g.DELETE("", product.Delete)
	}
}
