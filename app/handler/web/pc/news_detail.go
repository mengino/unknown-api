package pc

import (
	"dlsite/app/repository"
	"dlsite/internal/http/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewsDetail 新闻详情
func NewsDetail(c *gin.Context) {
	var u uri
	if err := c.BindUri(&u); err != nil {
		fmt.Println(err)
	}

	response.HTML(c, http.StatusOK, "news_detail.html", gin.H{
		"nav":  "news",
		"news": repository.GetArticle(u.ID),
	})
}
