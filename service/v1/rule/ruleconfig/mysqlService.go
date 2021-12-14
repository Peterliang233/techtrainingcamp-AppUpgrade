package ruleconfig

import (
	"net/http"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
)

// CreateRule 在mysql里面创建一条规则
func CreateRule(data *model.Rule) (int, int) {
	if err := mysql.Db.Create(data).Error; err != nil {
		return http.StatusInternalServerError, errmsg.ErrCreateRule
	}

	return http.StatusOK, errmsg.Success
}

// GetRules 在数据库获取所有的更新的规则
func GetRules(pageNum, pageSize int) ([]model.Rule, int, error) {
	var data []model.Rule
	var total int
	if err := mysql.Db.
		Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&data).Count(&total).
		Error; err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

// DeleteRuleFromMysql 从数据库里面删除这条规则
func DeleteRuleFromMysql(ruleID string) (int, int) {
	if err := mysql.Db.Where("id = ?", ruleID).Delete(&model.Rule{}).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}

	return http.StatusOK, errmsg.Success
}

// MysqlRuleOffline 从mysql里面的规则状态表格里设置为下线
func MysqlRuleOffline(id string) error {
	if err := mysql.Db.Model(&model.Rule{}).
		Where("id = ?", id).
		Update("status", false).
		Error; err != nil {
		return err
	}

	return nil
}

// MysqlRuleOnline 将数据库里面的这条规则的状态设置为上线
func MysqlRuleOnline(id string) error {
	if err := mysql.Db.Model(&model.Rule{}).
		Where("id = ?", id).
		Update("status", true).
		Error; err != nil {
		return err
	}

	return nil
}

// GetAllRules 在数据库里面获取所有的规则
func GetAllRules() ([]model.Rule, error) {
	var rules []model.Rule
	if err := mysql.Db.Find(&rules).Error; err != nil {
		return nil, err
	}
	return rules, nil
}

// GetDeviceWhiteList 获取规则对应的设备的白名单
func GetDeviceWhiteList() ([]model.Device, error) {
	var devices []model.Device
	if err := mysql.Db.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}
