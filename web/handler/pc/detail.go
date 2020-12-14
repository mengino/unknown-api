package pc

import (
	"dlsite/internal/http/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Detail App详情
func Detail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "detail.html", gin.H{
		"nav": "game",
		"app": app{
			ID:        3,
			Title:     "zzzzzz",
			Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
			Desc:      "浓浓的古典气息",
			Version:   "v1.0",
			Size:      "63MB",
			URL:       "https://aaa.com",
			Language:  "简体中文",
			CreatedAt: time.Now(),
			Category: category{
				ID:    1,
				Title: "角色扮演",
			},
		},
	})
}
