package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// News 新闻攻略
func News(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news.html", gin.H{
		"nav": "news",
	})
}
