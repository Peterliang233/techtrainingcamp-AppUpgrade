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
func CacheBasicInfo(platform, channelNumber string, cpuArch, appID, id int) {
	data := map[string]interface{}{
		"platform":       platform,
		"channel_number": channelNumber,
		"cpu_arch":       cpuArch,
		"app_id":         appID,
	}

	key := utils.MapToString(data)

	redis.RedisClient.SAdd(context.Background(), key, id)
}

// CacheUpdateVersionCode 将版本更新信息放到缓存里面
func CacheUpdateVersionCode(minUpdateVersionCode, maxUpdateVersionCode string, id int) {
	key := "app_update_version_code_" + strconv.Itoa(id)
	val := minUpdateVersionCode + ":" + maxUpdateVersionCode
	redis.RedisClient.Set(context.Background(), key, val, 0)
}

// CacheOsApi 将app_os_api放到缓存里面
func CacheOsApi(minOsApi, maxOsApi, id int) {
	key := "app_os_api_" + strconv.Itoa(id)
	val := strconv.Itoa(minOsApi) + ":" + strconv.Itoa(maxOsApi)
	redis.RedisClient.Set(context.Background(), key, val, 0)
}

// CacheOnline 将上线的规则id存放进缓存里面,这是一个集合，值为上线规则的id
func CacheOnline(id int) {
	redis.RedisClient.SAdd(context.Background(), "online", id, 0)
}

// CreateRule 在mysql里面创建一条规则
func CreateRule(data *model.Rule) (int, int) {
	if err := mysql.Db.Create(data).Error; err != nil {
		return http.StatusInternalServerError, errmsg.ErrCreateRule
	}
	// 在规则状态表里面插入一条数据,默认为true
	ruleState := &model.RuleState{
		RuleID: data.ID,
		State:  true,
	}
	if err := mysql.Db.Create(ruleState).Error; err != nil {
		return http.StatusInternalServerError, errmsg.ErrCreateRuleState
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
