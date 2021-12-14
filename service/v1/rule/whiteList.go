package rule

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
)

// AddWhiteList 添加某一条规则的白名单
func AddWhiteList(data *model.Device) (int, int) {
	if err := mysql.Db.
		Create(data).
		Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}

	return http.StatusOK, errmsg.Success
}

// CacheDeviceID 将设备白名单放到缓存里面
func CacheDeviceID(data *model.Device) {
	key := "app_device_id_" + strconv.Itoa(data.RuleID)
	redis.RedisClient.SAdd(context.Background(), key, data.DeviceID)
	redis.RedisClient.Expire(context.Background(), key, config.RedisSetting.ExpireTime)
}

// CheckDeviceIDFromWhiteList 检查白名单里面是否有这个设备ID
func CheckDeviceIDFromWhiteList(deviceID, ID string) bool {
	return redis.RedisClient.SIsMember(context.Background(), "app_device_id_"+ID, deviceID).Val()
}

// DeleteDeviceIDFromWhiteList 从白名单里面删除这个设备ID
func DeleteDeviceIDFromWhiteList(deviceID, ID string) error {
	// 先从redis里面删除这条ID
	redis.RedisClient.SRem(context.Background(), "app_device_id_"+ID, deviceID)
	// 然后从数据库里面删除
	return DeleteDeviceIDFromMysql(deviceID, ID)
}

// DeleteDeviceIDFromMysql 从数据库里面删除这条数据
func DeleteDeviceIDFromMysql(deviceID, ID string) error {
	if err := mysql.Db.
		Where("rule_id = ? AND device_id = ?", ID, deviceID).Delete(&model.Device{}).
		Error; err != nil {
		return err
	}

	return nil
}

// GetWhiteList 获取白名单列表
func GetWhiteList(pageNum, pageSize int) ([]model.Device, int, error) {
	var data []model.Device
	var total int
	if err := mysql.Db.Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Find(&data).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
