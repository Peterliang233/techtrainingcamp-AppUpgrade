package rule

import (
	"net/http"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/rule/ruleconfig"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	"github.com/gin-gonic/gin"
)

// ConfigRule 新版本规则配置接口
func ConfigRule(c *gin.Context) {
	var data model.Rule

	err := c.ShouldBindJSON(&data)
	data.Status = true
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": "参数请求错误",
				"data":   data,
			},
		})
		return
	}
	// 对这条规则持久化
	if statusCode, code := ruleconfig.CreateRule(&data); code != errmsg.Success {
		c.JSON(statusCode, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
				"data":   data,
			},
		})
		return
	}

	// 将需要的数据存入redis缓存里面
	code := ruleconfig.CacheOnline(strconv.Itoa(data.ID))
	code = ruleconfig.CacheBasicInfo(data.Platform, data.ChannelNumber, data.CPUArch, data.AppID, data.ID)
	code = ruleconfig.CacheOsApi(data.MinOSApi, data.MaxOSApi, data.ID)
	code = ruleconfig.CacheUpdateVersionCode(data.MinUpdateVersionCode, data.MaxUpdateVersionCode, data.ID)
	if code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[code],
				"data":   data,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": "新规则配置成功",
			"data":   data,
		},
	})
}

// GetRules 获取所有更新的规则
func GetRules(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	rules, total, err := ruleconfig.GetRules(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": "查询失败",
				"data":   nil,
				"total":  total,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.Success,
		"msg": map[string]interface{}{
			"detail": "查询成功",
			"data":   rules,
			"total":  total,
		},
	})
}

// DelRule 删除配置规则
func DelRule(c *gin.Context) {
	ruleID := c.Query("rule_id")

	statusCode, code := ruleconfig.DeleteRuleFromMysql(ruleID)

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data":   ruleID,
		},
	})
}

// OfflineRule 对这条规则执行下线处理
func OfflineRule(c *gin.Context) {
	id := c.Query("rule_id")

	// 先将这个id从redis里面删除
	if code := ruleconfig.RedisRuleOffline(id); code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.ErrOfflineRule,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrOfflineRule],
				"data":   id,
			},
		})
		return
	}

	// 从mysql里面将这个状态改为下线状态
	if err := ruleconfig.MysqlRuleOffline(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.ErrOfflineRule,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrOfflineRule],
				"data":   id,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.Success,
		"msg": map[string]interface{}{
			"detail": "下线成功",
			"data":   id,
		},
	})
}

// OnlineRule 规则上线
func OnlineRule(c *gin.Context) {
	id := c.Query("rule_id")

	// 先将这个id添加到redis里面
	if code := ruleconfig.CacheOnline(id); code != errmsg.Success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.ErrOnlineRule,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrOnlineRule],
				"data":   id,
			},
		})
		return
	}

	// 从mysql里面将这个状态改为上线状态
	if err := ruleconfig.MysqlRuleOnline(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.ErrOnlineRule,
			"msg": map[string]interface{}{
				"detail": errmsg.CodeMsg[errmsg.ErrOnlineRule],
				"data":   id,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.Success,
		"msg": map[string]interface{}{
			"detail": "上线成功",
			"data":   id,
		},
	})
}
