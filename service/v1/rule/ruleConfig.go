package rule

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/mysql"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
)

// CacheBasicInfo 将基本的信息存放到缓存里面
func CacheBasicInfo(platform, channelNumber, cpuArch string, AID, id int) {
	data := map[string]interface{}{
		"platform":       platform,
		"channel_number": channelNumber,
		"cpu_arch":       cpuArch,
		"aid":            AID,
	}

	key := utils.MapToString(data)

	redis.RedisClient.SAdd(context.Background(), key, id)
}

// CacheUpdateVersionCode 将版本更新信息放到缓存里面
func CacheUpdateVersionCode(minUpdateVersionCode, maxUpdateVersionCode string, id int) {
	key := "app_update_version_code_" + strconv.Itoa(id)
	val := minUpdateVersionCode + ":" + maxUpdateVersionCode
	redis.RedisClient.SAdd(context.Background(), key, val)
}

// CacheOsApi 将app_os_api放到缓存里面
func CacheOsApi(minOsApi, maxOsApi, id int) {
	key := "app_os_api_" + strconv.Itoa(id)
	val := strconv.Itoa(maxOsApi) + ":" + strconv.Itoa(minOsApi)
	redis.RedisClient.SAdd(context.Background(), key, val)
}

// CreateRule 在mysql里面创建一条规则
func CreateRule(data *model.Rule) (int, int) {
	if err := mysql.Db.Create(data).Error; err != nil {
		return http.StatusInternalServerError, errmsg.Error
	}
	return http.StatusOK, errmsg.Success
}
