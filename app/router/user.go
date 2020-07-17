package router

import (
	"dlsite/app/handler/user"

	"github.com/gin-gonic/gin"
)

func userGroup(r *gin.RouterGroup) {
	g := r.Group("user")
	{
		g.GET("current", user.Current)
	}
}
