package pc

import (
	"dlsite/internal/http/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Top 排行榜
func Top(c *gin.Context) {
	response.HTML(c, http.StatusOK, "top.html", gin.H{
		"nav": "top",
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
		"category": []category{
			{
				ID:   1,
				Name: "角色扮演",
			},
			{
				ID:   2,
				Name: "休闲益智",
			},
			{
				ID:   3,
				Name: "动作冒险",
			},
			{
				ID:   4,
				Name: "解谜闯关",
			},
		},
	})
}
