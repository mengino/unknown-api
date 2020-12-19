package pc

import (
	"dlsite/app/model"
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
			ID:       3,
			Title:    "zzzzzz",
			Image:    "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
			Desc:     "浓浓的古典气息",
			Version:  "v1.0",
			Size:     "63MB",
			URL:      "https://aaa.com",
			Language: "简体中文",
			Content:  "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
			Slide: model.ProductSlideJSON{
				"https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				"https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				"https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				"https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				"https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
			},
			CreatedAt: time.Now(),
			Category: category{
				ID:   1,
				Name: "角色扮演",
			},
		},
	})
}
