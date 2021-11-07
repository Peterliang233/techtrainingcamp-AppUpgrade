package utils

import (
	"time"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/golang-jwt/jwt"
)

var MySecret = []byte("BYteDance")

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, int) {
	Claims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "Peterliang",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodES256, Claims)

	token, err := reqClaims.SignedString(MySecret)

	if err != nil {
		return "", errmsg.Error
	}

	return token, errmsg.Success
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
