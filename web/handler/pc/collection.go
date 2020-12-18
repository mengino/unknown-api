package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Collection 合集
func Collection(c *gin.Context) {
	response.HTML(c, http.StatusOK, "collection.html", gin.H{
		"nav": "hj",
	})
}
