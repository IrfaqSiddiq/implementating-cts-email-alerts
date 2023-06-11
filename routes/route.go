package routes

import (
	"cts-alerts/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.RouterGroup) {
	api := router.Group("/api")
	{
		api.POST("/send-email-alerts",controllers.SendEmailAlerts)
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode("debug")
	AddRoutes(&router.RouterGroup)
	return router
}
