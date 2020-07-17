package middleware

import (
	"dlsite/internal/auth"
	"dlsite/internal/http/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Validate 验证JWT Token
func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Exception(c, http.StatusUnprocessableEntity, "请求头中auth为空")
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Exception(c, http.StatusUnprocessableEntity, "请求头中auth格式有误")
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := auth.ParseToken(parts[1])
		if err != nil {
			response.Exception(c, http.StatusBadRequest, "无效的Token")
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("uid", int(mc.UID))
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
