package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HjDetail 合集详情
func HjDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "hj_detail.html", gin.H{
		"nav": "hj",
	})
}
