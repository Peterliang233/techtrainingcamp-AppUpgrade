package rule

import (
	"net/http"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	ruleService "github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/rule"
	"github.com/gin-gonic/gin"
)

// RuleConfig 新版本规则配置接口
func RuleConfig(c *gin.Context) {
	var data model.Rule

	err := c.ShouldBindJSON(&data)
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
	statusCode, code := ruleService.CreateRule(&data)
	// 将需要的数据存入redis缓存里面
	ruleService.CacheOnline(data.ID)
	ruleService.CacheBasicInfo(data.Platform, data.ChannelNumber, data.CPUArch, data.AppID, data.ID)
	ruleService.CacheOsApi(data.MinOSApi, data.MaxOSApi, data.ID)
	ruleService.CacheUpdateVersionCode(data.MinUpdateVersionCode, data.MaxUpdateVersionCode, data.ID)
	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data":   data,
		},
	})
}

// GetRules 获取所有更新的规则
func GetRules(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	rules, total, err := ruleService.GetRules(pageNum, pageSize)
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

	statusCode, code := ruleService.DeleteRuleFromMysql(ruleID)

	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data":   ruleID,
		},
	})
}
