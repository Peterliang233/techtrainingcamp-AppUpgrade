package ruleconfig

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/database/redis"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/utils"
)

// CacheBasicInfo 将基本的信息存放到缓存里面
func CacheBasicInfo(platform, channelNumber string, cpuArch, appID, id int) int {
	data := map[string]interface{}{
		"platform":       platform,
		"channel_number": channelNumber,
		"cpu_arch":       cpuArch,
		"app_id":         appID,
	}

	key := utils.MapToString(data)

	_, err := redis.RedisClient.SAdd(context.Background(), key, id).Result()
	if err != nil {
		return errmsg.ErrCacheBasicInfo
	}

	return errmsg.Success
}

// CacheUpdateVersionCode 将版本更新信息放到缓存里面
func CacheUpdateVersionCode(minUpdateVersionCode, maxUpdateVersionCode string, id int) int {
	key := "app_update_version_code_" + strconv.Itoa(id)
	val := minUpdateVersionCode + ":" + maxUpdateVersionCode
	_, err := redis.RedisClient.Set(context.Background(), key, val, config.RedisSetting.ExpireTime).Result()
	if err != nil {
		return errmsg.ErrCacheUpdateVersionCode
	}

	return errmsg.Success
}

// CacheOsApi 将app_os_api放到缓存里面
func CacheOsApi(minOsApi, maxOsApi, id int) int {
	key := "app_os_api_" + strconv.Itoa(id)
	val := strconv.Itoa(minOsApi) + ":" + strconv.Itoa(maxOsApi)
	_, err := redis.RedisClient.Set(context.Background(), key, val, config.RedisSetting.ExpireTime).Result()
	if err != nil {
		return errmsg.ErrCacheApi
	}
	return errmsg.Success
}

// CacheOnline 将上线的规则id存放进缓存里面,这是一个集合，key为online，val为上线规则的id
func CacheOnline(id string) int {
	_, err := redis.RedisClient.SAdd(context.Background(), "online", id).Result()
	if err != nil {
		fmt.Println(err)
		return errmsg.ErrOnlineRule
	}
	return errmsg.Success
}

// RedisRuleOffline 从redis的online的集合里面删除这个id
func RedisRuleOffline(id string) int {
	_, err := redis.RedisClient.SRem(context.Background(), "online", id).Result()
	if err != nil {
		return errmsg.ErrOfflineRule
	}

	return errmsg.Success
}

// RuleOffline 检查规则是否下线
func RuleOffline(id string) bool {
	return redis.RedisClient.SIsMember(context.Background(), "online", id).Val()
}
