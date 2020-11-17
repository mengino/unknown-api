package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	response.HTML(c, http.StatusOK, "pc/index.html", "pc")
}
