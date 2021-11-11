package rule

import (
	"context"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
)

// CacheBasicInfo 将基本的信息存放到缓存里面
func CacheBasicInfo(platform, channelNumber, cpuArch string, api, id int) {
	data := map[string]interface{}{
		"platform":       platform,
		"channel_number": channelNumber,
		"cpu_arch":       cpuArch,
		"api":            api,
	}

	key := utils.MapToString(data)

	redis.RedisClient.SAdd(context.Background(), key, id)
}

//CacheUpdateVersionCode 将版本更新信息放到缓存里面
func CacheUpdateVersionCode(minUpdateVersionCode, maxUpdateVersionCode string, id int) {
	key := "app_version_" + strconv.Itoa(id)
	val := minUpdateVersionCode + ":" + maxUpdateVersionCode
	redis.RedisClient.SAdd(context.Background(), key, val)
}

// CacheOsApi 将app_os_api放到缓存里面
func CacheOsApi(minOsApi, maxOsApi string, id int) {
	key := "app_os_api_" + strconv.Itoa(id)
	val := minOsApi + ":" + maxOsApi
	redis.RedisClient.SAdd(context.Background(), key, val)
}

// CacheDeviceID 将设备白名单放到缓存里面
func CacheDeviceID(deviceID string, id int) {
	key := "app_device_id_" + strconv.Itoa(id)
	redis.RedisClient.SAdd(context.Background(), key, deviceID)
}
