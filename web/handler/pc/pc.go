package pc

import (
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 首页
func Index(c *gin.Context) {
	response.HTML(c, http.StatusOK, "index.html", "pc")
}

// Soft 首页
func Soft(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game.html", "pc")
}

// Game 首页
func Game(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game.html", "pc")
}

// Detail 首页
func Detail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "game_detail.html", "pc")
}

// News 首页
func News(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news.html", "pc")
}

// NewsDetail 首页
func NewsDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "news_detail.html", "pc")
}

// Hj 首页
func Hj(c *gin.Context) {
	response.HTML(c, http.StatusOK, "hj.html", "pc")
}

// HjDetail 首页
func HjDetail(c *gin.Context) {
	response.HTML(c, http.StatusOK, "hj_detail.html", "pc")
}

// Top 首页
func Top(c *gin.Context) {
	response.HTML(c, http.StatusOK, "top.html", "pc")
}
