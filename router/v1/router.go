package v1

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
	ruleApi "github.com/Peterliang233/techtrainingcamp-AppUpgrade/router/v1/api/rule"
	userApi "github.com/Peterliang233/techtrainingcamp-AppUpgrade/router/v1/api/user"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.ServerSetting.AppMode)

	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	api := r.Group("/api")

	// 用户接口模块
	user := api.Group("/user")
	user.POST("/sign_in", userApi.SignIn)
	user.Use(middleware.JWTAuthMiddleware())
	{
		user.POST("/sign_up", userApi.SignUp)
	}

	// 规则配置接口模块
	rule := api.Group("/rule")
	rule.Use(middleware.JWTAuthMiddleware())
	{
		rule.POST("/settings", ruleApi.RuleConfig)
		rule.GET("/all", ruleApi.GetRules)
		rule.POST("/verification", ruleApi.RuleCheck)
		rule.POST("/whitelist", ruleApi.AddDeviceID)
	}

	return r
}
