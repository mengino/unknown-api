package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hj 合集
func Hj(c *gin.Context) {
	response.HTML(c, http.StatusOK, "hj.html", gin.H{
		"nav": "hj",
	})
}
