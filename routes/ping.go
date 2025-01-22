package routes

import (
	"AgentApiGo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPingRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
}
