package middleware

import (
	"net/http"
	"strings"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 鉴权认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": errmsg.Error,
				"data": map[string]interface{}{
					"status": errmsg.CodeMsg[errmsg.AuthEmpty],
				},
			})
			c.Abort()

			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": errmsg.Error,
				"data": map[string]interface{}{
					"status": errmsg.InvalidToken,
					"msg":    errmsg.CodeMsg[errmsg.InvalidToken],
				},
			})
			c.Abort()

			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": errmsg.Error,
				"data": map[string]interface{}{
					"status": errmsg.InvalidToken,
					"msg":    err.Error(),
				},
			})
			c.Abort()

			return
		}

		c.Set("username", claims.Username)

		c.Next()
	}
}
