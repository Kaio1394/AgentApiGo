package routes

import (
	"AgentApiGo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublishRoutes(r *gin.Engine) {
	r.POST("/publish", controllers.Publish)
}
