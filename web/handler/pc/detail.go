package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Detail App详情
func Detail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game_detail.html", gin.H{
		"nav": "game",
	})
}
