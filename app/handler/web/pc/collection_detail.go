package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CollectionDetail 合集详情
func CollectionDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "collection_detail.html", gin.H{
		"nav": "hj",
	})
}
