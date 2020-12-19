package pc

import (
	"dlsite/app/model"
	"dlsite/app/repository"
	"dlsite/internal/http/response"
	"net/http"
	"strconv"
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

type query struct {
	Category int
	Page     int
	Sort     int
}

// Game 游戏下载
func Game(c *gin.Context) {
	category, err := strconv.Atoi(c.Query("category"))
	if err != nil {

	}

	response.HTML(c, http.StatusOK, "app.html", gin.H{
		"nav":      "game",
		"category": repository.GetCategoryList(model.GroupGame),
		"list":     repository.GetProductList(model.GroupGame, category, c.DefaultQuery("sort", "sort")),
	})
}

// Soft 软件下载
func Soft(c *gin.Context) {
	response.HTML(c, http.StatusOK, "app.html", gin.H{
		"nav":      "soft",
		"category": repository.GetCategoryList(model.GroupSoft),
		"list": []app{
			{
				ID:        1,
				Title:     "微信信息发出去对方没收到",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "浓浓的古典气息",
				Version:   "v1.0",
				Size:      "63MB",
				CreatedAt: time.Now(),
				Category: category{
					ID:   1,
					Name: "角色扮演",
				},
			},
			{
				ID:        2,
				Title:     "xxxxxx",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "浓浓的古典气息",
				Version:   "v1.0",
				Size:      "63MB",
				CreatedAt: time.Now(),
				Category: category{
					ID:   1,
					Name: "角色扮演",
				},
			},
			{
				ID:        3,
				Title:     "zzzzzz",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "浓浓的古典气息",
				Version:   "v1.0",
				Size:      "63MB",
				CreatedAt: time.Now(),
				Category: category{
					ID:   1,
					Name: "角色扮演",
				},
			},
			{
				ID:        4,
				Title:     "bbbbbs",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "浓浓的古典气息",
				Version:   "v1.0",
				Size:      "63MB",
				CreatedAt: time.Now(),
				Category: category{
					ID:   1,
					Name: "角色扮演",
				},
			},
			{
				ID:        5,
				Title:     "aaaa",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "浓浓的古典气息",
				Version:   "v1.0",
				Size:      "63MB",
				CreatedAt: time.Now(),
				Category: category{
					ID:   1,
					Name: "角色扮演",
				},
			},
		},
	})
}
