package user

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
	"github.com/jinzhu/gorm"
)

// JudgeUsername 判断用户名，我们禁止用户名命名为admin和用户名重复
func JudgeUsername(username string) bool {
	if username == "admin" {
		return false
	}

	// 判断用户名是否存在
	if err := mysql.Db.
		Where("username = ?", username).
		First(&model.User{}).
		Error; err == gorm.ErrRecordNotFound {
		return true
	}

	return false
}

// CreateUser 创建新的用户
func CreateUser(user *model.User) int {
	user.Password = utils.EncryptPassword(user.Password)
	if err := mysql.Db.Create(user).Error; err != nil {
		return errmsg.Error
	}

	return errmsg.Success
}
