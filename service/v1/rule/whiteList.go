package rule

import (
	"context"
	"net/http"
	"strconv"

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
}
