package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewsDetail 新闻详情
func NewsDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news_detail.html", gin.H{
		"nav": "news",
	})
}
