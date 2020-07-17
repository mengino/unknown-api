package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success 查询成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code":    0,
			"data":    data,
			"message": "操作成功",
		},
	)
}

// SuccessCreated 创建/修改成功返回
func SuccessCreated(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusCreated,
		map[string]interface{}{
			"code":    0,
			"data":    data,
			"message": "操作成功",
		},
	)
}

// SuccessNoContent 删除成功返回
func SuccessNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

// Fail 失败返回
func Fail(c *gin.Context, code int, message string) {
	c.JSON(
		code,
		map[string]interface{}{
			"code":    code,
			"data":    struct{}{},
			"message": message,
		},
	)
}

// Exception 异常返回
func Exception(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, map[string]interface{}{
		"code":    code,
		"data":    struct{}{},
		"message": message,
	})
}
