package pc

import (
	"dlsite/internal/http/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// News 新闻攻略
func News(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news.html", gin.H{
		"nav": "news",
		"news": []news{
			{
				ID:        1,
				Title:     "微信信息发出去对方没收到",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Title:     "xxxxxx",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
			},
			{
				ID:        3,
				Title:     "zzzzzz",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
			},
			{
				ID:        4,
				Title:     "bbbbbs",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
			},
			{
				ID:        5,
				Title:     "aaaa",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
			},
		},
		"category": []category{
			{
				ID:    1,
				Title: "新闻资讯",
			},
			{
				ID:    2,
				Title: "游戏攻略",
			},
			{
				ID:    3,
				Title: "软件教程",
			},
			{
				ID:    4,
				Title: "活动大全",
			},
		},
	})
}
