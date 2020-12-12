package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Game 游戏下载
func Game(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game.html", gin.H{
		"nav": "game",
	})
}
