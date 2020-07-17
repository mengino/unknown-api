package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims 自定义Token结构
type Claims struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}

const tokenExpireDuration = time.Hour * 24

var secret = []byte("1HdHa99mxSmavaVddehGBQaccdaB1b")

// GenToken 创建token
func GenToken(uid int) (string, error) {
	// 创建一个我们自己的声明
	c := Claims{
		uid, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(), // 过期时间
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
