package v1

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/config"
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

	api.POST("/sign_in", userApi.SignIn)

	api.Use(middleware.JWTAuthMiddleware())
	{
		api.POST("/sign_up", userApi.SignUp)
	}

	return r
}
