package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Soft 软件下载
func Soft(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game.html", gin.H{
		"nav": "soft",
	})
}
