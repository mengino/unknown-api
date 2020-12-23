package pc

import (
	"dlsite/app/model"
	"dlsite/app/repository"
	"dlsite/internal/http/response"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type app struct {
	ID        int
	Title     string
	Image     string
	Desc      string
	Version   string
	Size      string
	URL       string
	Language  string
	Content   string
	Slide     model.ProductSlideJSON
	CreatedAt time.Time
	Category  category
}

// Game 游戏下载
func Game(c *gin.Context) {
	var q query
	if err := c.BindQuery(&q); err != nil {
		fmt.Println(err)
	}

	response.HTML(c, http.StatusOK, "app.html", gin.H{
		"nav":      "game",
		"params":   q,
		"category": repository.GetCategoryList(model.GroupGame),
		"list":     repository.GetProductList(model.GroupGame, q.Category, q.Sort),
	})
}

// Soft 软件下载
func Soft(c *gin.Context) {
	var q query
	if err := c.BindQuery(&q); err != nil {
		fmt.Println(err)
	}

	response.HTML(c, http.StatusOK, "app.html", gin.H{
		"nav":      "soft",
		"params":   q,
		"category": repository.GetCategoryList(model.GroupSoft),
		"list":     repository.GetProductList(model.GroupSoft, q.Category, q.Sort),
	})
}
