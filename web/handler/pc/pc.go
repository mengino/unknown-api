package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	response.HTML(c, http.StatusOK, "index.html", gin.H{
		"nav": "index",
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
