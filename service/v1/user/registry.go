package user

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
)

// JudgeUsername 判断用户名，我们禁止用户名命名为admin
func JudgeUsername(username string) bool {
	return username == "admin"
}

// CreateUser 创建新的用户
func CreateUser(user *model.User) int {
	user.Password = utils.EncryptPassword(user.Password)
	if err := mysql.Db.Create(user).Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}
