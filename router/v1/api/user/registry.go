package user

import (
	"net/http"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"

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

	msg, code := utils.Validate(user)

	if code != errmsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": msg,
				"data":   nil,
			},
		})
		return
	}

	// 如果用户名为admin，则禁止注册该帐号
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

	// 设定只能admin账户才能进行注册
	name, _ := c.Get("username")

	if name != config.AdminSetting.Username {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": "非管理员用户不能进行注册",
				"data":   nil,
			},
		})
		return
	}

	code = userService.CreateUser(&user)

	if code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": "服务端错误，注册失败",
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
