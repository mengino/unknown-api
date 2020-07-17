package user

import (
	"dlsite/app/model"
	"dlsite/internal/auth"
	"dlsite/internal/db"
	"dlsite/internal/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login 绑定 JSON
type loginRequest struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token            string `json:"token"`
	CurrentAuthority string `json:"current_authority"`
}

// Login 登录接口
func Login(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user model.User

	if db.New().Where("nickname = ?", request.Nickname).First(&user).RecordNotFound() {
		response.Fail(c, http.StatusBadRequest, "登录信息有误，请核对后再试")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		response.Fail(c, http.StatusUnauthorized, "登录信息有误，请核对后再试")
		return
	}

	if token, err := auth.GenToken(user.ID); err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, &loginResponse{Token: token, CurrentAuthority: "admin"})
	}
}

type currentUserResponse struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

// Current 获取当前用户
func Current(c *gin.Context) {
	var user model.User

	db.New().Where("id = ?", c.GetInt("uid")).First(&user)

	response.Success(c, &currentUserResponse{
		Avatar: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Name:   user.Nickname,
		UserID: int(user.ID),
	})
}
