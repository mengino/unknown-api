package pc

import (
	"dlsite/internal/http/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NewsDetail 新闻详情
func NewsDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news_detail.html", gin.H{
		"nav": "news",
		"news": news{
			ID:        4,
			CreatedAt: time.Now(),
			Title:     "bbbbbs",
			Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
			Content:   "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
			Author:    "xxxxx",
			Category: category{
				ID:    1,
				Title: "角色扮演",
			},
			Product: app{
				ID:        1,
				Title:     "微信信息发出去对方没收到",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "浓浓的古典气息",
				Version:   "v1.0",
				Size:      "63MB",
				CreatedAt: time.Now(),
			},
		},
	})
}
