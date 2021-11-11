package rule

import (
	"context"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
)

// GetUpdateInfo 根据客户端上报的参数返回对应的下载包信息
func GetUpdateInfo(info *model.Info) (*model.Rule, int) {
	var data []*model.Rule
	if err := mysql.Db.
		Where(map[string]interface{}{
			"aid":            info.AID,
			"channel_number": info.ChannelNumber,
			"cpu_arch":       info.CPUArch,
			"platform":       info.DevicePlatform,
		}).
		Where("min_update_version_code <= ?", info.UpdateVersionCode).
		Where("max_update_version_code >= ?", info.UpdateVersionCode).
		Where("min_os_api <= ?", info.OSApi).
		Where("max_os_api >= ?", info.OSApi).
		Find(&data).
		Error; err != nil {
		return nil, errmsg.Error
	}

	// 循环遍历整个数组，找到第一个满足请求参数的device_id在设备白名单里面
	for _, rule := range data {
		if ExistsDeviceId(strconv.Itoa(rule.ID), info.DeviceID) {
			return rule, errmsg.Success
		}
	}

	return nil, errmsg.Error
}

// ExistsDeviceId 检查redis里面这个id对应的集合里面是否有这个, 一个id对应的集合里面存储了这条规则对应的白名单
func ExistsDeviceId(id, deviceID string) bool {
	ok, _ := redis.RedisClient.SIsMember(context.Background(), id, deviceID).Result()
	return ok
}
