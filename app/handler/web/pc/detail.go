package pc

import (
	"dlsite/app/repository"
	"dlsite/internal/http/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Detail App详情
func Detail(c *gin.Context) {
	var u uri
	if err := c.BindUri(&u); err != nil {
		fmt.Println(err)
	}

	response.HTML(c, http.StatusOK, "detail.html", gin.H{
		"nav": "game",
		"app": repository.GetProduct(u.ID),
	})
}
