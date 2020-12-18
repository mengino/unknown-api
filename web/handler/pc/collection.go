package pc

import (
	"dlsite/internal/http/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type collection struct {
	ID        int
	Title     string
	Image     string
	Desc      string
	Category  int
	App       []app
	CreatedAt time.Time
}

// Collection 合集
func Collection(c *gin.Context) {
	response.HTML(c, http.StatusOK, "collection.html", gin.H{
		"nav": "hj",
	})
}
