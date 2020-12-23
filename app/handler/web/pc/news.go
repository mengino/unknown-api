package pc

import (
	"dlsite/app/model"
	"dlsite/app/repository"
	"dlsite/internal/http/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// News 新闻攻略
func News(c *gin.Context) {
	var q query
	if err := c.BindQuery(&q); err != nil {
		fmt.Println(err)
	}

	response.HTML(c, http.StatusOK, "news.html", gin.H{
		"nav":      "news",
		"params":   q,
		"category": repository.GetCategoryList(model.GroupNews),
		"news":     repository.GetArticleList(q.Category),
	})
}
