package user

import (
	"net/http"

	userService "github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/user"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	"github.com/gin-gonic/gin"
)

// SignUp 注册接口
func SignUp(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.Error],
				"data":   user,
			},
		})
		return
	}

	// 如果用户名为admin，则禁止注册
	if !userService.JudgeUsername(user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.ErrUsername,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrUsername],
				"data":   user,
			},
		})
		return
	}

	// 判断是否有重复的用户名

	code := userService.CreateUser(&user)

	if code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": "注册失败",
				"data":   user.Username,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": "注册成功",
			"data":   user.Username,
		},
	})
}
