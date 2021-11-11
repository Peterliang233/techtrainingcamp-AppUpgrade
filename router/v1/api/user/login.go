package user

import (
	"net/http"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	userService "github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/user"
	"github.com/gin-gonic/gin"
)

// SignIn 登录接口
func SignIn(c *gin.Context) {
	var login model.User
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.Error],
				"data":   login,
			},
		})
		return
	}

	code := userService.CheckLogin(login.Username, login.Password)
	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": "登录失败",
				"data":   nil,
			},
		})
		return
	}

	var tokenString string
	tokenString, code = utils.GenerateToken(login.Username)
	if code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
				"data":   nil,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": "登录成功",
			"token":  tokenString,
		},
	})
}
