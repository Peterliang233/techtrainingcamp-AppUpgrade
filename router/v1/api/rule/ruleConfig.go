package rule

import (
	"net/http"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/model"
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

}
