package utils

import (
	"encoding/base64"
	"log"
	"time"

	"golang.org/x/crypto/scrypt"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/dgrijalva/jwt-go"
)

var MySecret = []byte("ByteDance")

type MyClaims struct {
	Username string `json:"username"` // 利用中间件保存一些有用的信息
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(username string) (string, int) {
	Claims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // 设置过期时间
			Issuer:    "Peterliang",                         // 设置签发人
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

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

// EncryptPassword 密码加密
func EncryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{23, 32, 21, 11, 11, 22, 11, 0}
	HashPassword, err := scrypt.Key([]byte(password), salt, 32768, 6, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(HashPassword)
}
