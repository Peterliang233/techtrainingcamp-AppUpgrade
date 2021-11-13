package rule

import (
	"net/http"
	"strconv"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
	ruleService "github.com/Peterliang233/techtrainingcamp-AppUpgrade/service/v1/rule"
	"github.com/gin-gonic/gin"
)

// AddDeviceID 给具体id的规则添加白名单
func AddDeviceID(c *gin.Context) {
	var data model.Device
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
	// 白名单存入redis
	ruleService.CacheDeviceID(&data)
	// 白名单持久化
	statusCode, code := ruleService.AddWhiteList(&data)
	c.JSON(statusCode, gin.H{
		"code": code,
		"msg": map[string]interface{}{
			"detail": errmsg.CodeMsg[code],
			"data":   data,
		},
	})
}

// GetWhiteList 查找某一个device_id是否在某一条规则的白名单
func GetWhiteList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	data, count, err := ruleService.GetWhiteList(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": "获取白名单失败",
				"data":   data,
				"total":  count,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.Success,
		"msg": map[string]interface{}{
			"detail": "获取白名单成功",
			"data":   data,
			"total":  count,
		},
	})
}

// DelFromWhite 从某一条规则的白名单里面删除一条执行的deviceID接口
func DelFromWhiteList(c *gin.Context) {
	ruleID := c.Query("rule_id")
	deviceID := c.Query("device_id")

	// 先要检查这个设备ID是否在这个白名单里面
	if ok := ruleService.CheckDeviceIDFromWhiteList(deviceID, ruleID); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": "该设备ID不存在白名单里面",
				"data": map[string]interface{}{
					"device_id": deviceID,
					"rule_id":   ruleID,
				},
			},
		})
		return
	}

	err := ruleService.DeleteDeviceIDFromWhiteList(deviceID, ruleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errmsg.Error,
			"msg": map[string]interface{}{
				"detail": "删除失败",
				"data": map[string]interface{}{
					"device_id": deviceID,
					"rule_id":   ruleID,
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": errmsg.Success,
		"msg": map[string]interface{}{
			"detail": "删除成功",
			"data": map[string]interface{}{
				"device_id": deviceID,
				"rule_id":   ruleID,
			},
		},
	})
}
