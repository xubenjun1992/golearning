package utils

import (
	"blog/core"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64, userName string) string {
	claims := jwt.MapClaims{
		"userId":   userId,
		"userName": userName,
		"exp":      time.Now().Add(time.Duration(core.Configs.Timeout) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(core.Configs.Secret))
	if err != nil {
		panic("生成 JWT 失败: " + err.Error())
	}
	return tokenString
}

func ParseToken(tokenString string, options ...jwt.ParserOption) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(core.Configs.Secret), nil
	}, options...)
	if err != nil {
		return nil, err
	}
	return token, nil
}
