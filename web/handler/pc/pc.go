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
	CreatedAt time.Time
}
type category struct {
	ID        int
	Title     string
	CreatedAt time.Time
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
				ID:    1,
				Title: "角色扮演",
			},
			{
				ID:    2,
				Title: "休闲益智",
			},
			{
				ID:    3,
				Title: "动作冒险",
			},
			{
				ID:    4,
				Title: "解谜闯关",
			},
		},
	})
}

// Soft 软件下载
func Soft(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game.html", gin.H{
		"nav": "soft",
	})
}

// Game 游戏下载
func Game(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game.html", gin.H{
		"nav": "game",
	})
}

// Detail App详情
func Detail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game_detail.html", gin.H{
		"nav": "game",
	})
}

// News 新闻攻略
func News(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news.html", gin.H{
		"nav": "news",
	})
}

// NewsDetail 新闻详情
func NewsDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news_detail.html", gin.H{
		"nav": "news",
	})
}

// Hj 合集
func Hj(c *gin.Context) {
	response.HTML(c, http.StatusOK, "hj.html", gin.H{
		"nav": "hj",
	})
}

// HjDetail 合集详情
func HjDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "hj_detail.html", gin.H{
		"nav": "hj",
	})
}

// Top 排行榜
func Top(c *gin.Context) {
	response.HTML(c, http.StatusOK, "top.html", gin.H{
		"nav": "top",
	})
}
