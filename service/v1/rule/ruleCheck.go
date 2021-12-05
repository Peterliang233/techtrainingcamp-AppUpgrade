package rule

import (
	"context"
	"strconv"
	"strings"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
)

// GetUpdateInfo 根据客户端上报的参数返回对应的下载包信息
func GetUpdateInfo(info *model.Info) (*model.Rule, int) {
	basicID := GetBasicID(info.DevicePlatform,
		info.ChannelNumber,
		info.CPUArch,
		info.AppID)

	id, ok := GetRuleID(basicID, info.UpdateVersionCode, info.DeviceID, info.OSApi)
	if !ok {
		return nil, errmsg.Error
	}
	data := new(model.Rule)
	if err := mysql.Db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, errmsg.Error
	}
	return data, errmsg.Success
}

// 获取初步符合基本信息的ID集合
func GetBasicID(platform, channelNumber string, cpuArch, appID int) []string {
	data := map[string]interface{}{
		"platform":       platform,
		"channel_number": channelNumber,
		"cpu_arch":       cpuArch,
		"app_id":         appID,
	}

	key := utils.MapToString(data)
	return redis.RedisClient.SMembers(context.Background(), key).Val()
}

// GetRuleID 根据之前设置的redis，返回命中的规则的ID
func GetRuleID(ID []string, versionCode, deviceID string, oSApi int) (string, bool) {
	nowApi := strconv.Itoa(oSApi)
	for _, id := range ID {
		// 检查deviceID是否在redis里面这条规则的白名单集合里面
		if ok := redis.RedisClient.SIsMember(context.Background(),
			"app_device_id_"+id, deviceID).Val(); !ok {
			return "0", false
		}
		version := redis.RedisClient.Get(context.Background(),
			"app_update_version_code_"+id).Val()

		updateVersionCode := strings.Split(version, ":")

		if len(updateVersionCode) != 2 {
			continue
		}
		// 如果该版本在最大和最小更新版本之间，那么就可以进一步筛选
		if utils.CompareVersion(updateVersionCode[0], versionCode) <= 0 &&
			utils.CompareVersion(updateVersionCode[1], versionCode) >= 0 {
			// 检查os_api是否在一个合法的区间里面
			api := redis.RedisClient.Get(context.Background(),
				"app_os_api_"+id).Val()
			osApi := strings.Split(api, ":")

			if len(osApi) != 2 {
				continue
			}

			// 如果同时也满足在指定的os_api的版本之间
			if strings.Compare(osApi[0], nowApi) <= 0 &&
				strings.Compare(osApi[1], nowApi) >= 0 {
				return id, true
			}
		}
	}

	return "0", false
}
