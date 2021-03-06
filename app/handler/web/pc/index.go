package pc

import (
	"dlsite/internal/http/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type news struct {
	ID        int
	Title     string
	Image     string
	Desc      string
	Total     int
	Author    string
	Content   string
	CreatedAt time.Time

	Product  app
	Category category
}

type category struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

type query struct {
	Page     int    `form:"page,default=1"`
	Category int    `form:"category"`
	Sort     string `form:"sort,default=sort"`
}

type uri struct {
	ID int `uri:"id" binding:"required"`
}

// Index 首页
func Index(c *gin.Context) {
	response.HTML(c, http.StatusOK, "index.html", gin.H{
		"nav": "index",
		"news": []news{
			{
				ID:        1,
				Title:     "微信信息发出去对方没收到",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
				Total:     50,
			},
			{
				ID:        2,
				Title:     "xxxxxx",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
				Total:     50,
			},
			{
				ID:        3,
				Title:     "zzzzzz",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
				Total:     50,
			},
			{
				ID:        4,
				Title:     "bbbbbs",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
				Total:     50,
			},
			{
				ID:        5,
				Title:     "aaaa",
				Image:     "https://img.homevv.com/uploadimg/ico/2019/0605/1559728456211488.jpg",
				Desc:      "微信作为大家常用的一款聊天软件，基本上大部分人都在用吧~而就在刚在，突然发生了一件奇怪的事！自己发出去的信息，也没有被拒收，明明发送成功了，可是对方却没有收到！？这是怎么回事！天呐撸，下面就让小编来为大家介绍一下吧！",
				CreatedAt: time.Now(),
				Total:     50,
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
