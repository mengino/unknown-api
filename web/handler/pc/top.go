package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Top 排行榜
func Top(c *gin.Context) {
	response.HTML(c, http.StatusOK, "top.html", gin.H{
		"nav": "top",
	})
}
