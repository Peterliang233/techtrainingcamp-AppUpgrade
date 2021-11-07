package v1

import (
	"net/http"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "This is a test",
		})
	})

	api := r.Group("/api")

	api.Use(middleware.JWTAuthMiddleware())

	return r
}
