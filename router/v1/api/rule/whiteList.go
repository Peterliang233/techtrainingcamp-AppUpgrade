package rule

import (
	"net/http"

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

// GetWhiteList 根据ID获取对应的白名单
func GetWhiteList(c *gin.Context) {

}
