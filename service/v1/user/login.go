package user

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
)

// CheckLogin 检查是否有登录权限
func CheckLogin(username, password string) int {
	// 先判断是否为管理员用户
	if username == config.AdminSetting.Username &&
		password == config.AdminSetting.Password {
		return errmsg.Success
	}

	var login model.User
	if err := mysql.Db.Where("username = ?", username).First(&login).Error; err != nil {
		return errmsg.Error
	}

	if utils.EncryptPassword(password) != login.Password {
		return errmsg.ErrPassword
	}

	return errmsg.Success
}
