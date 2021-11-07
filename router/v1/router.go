package v1

import (
	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.JWTAuthMiddleware())
	r.Use(middleware.Logger())

	return r
}
