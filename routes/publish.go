package routes

import (
	"AgentApiGo/controllers"

	"github.com/gin-gonic/gin"
)

// Publish publishes a job message to RabbitMQ
// @Summary Publish a job
// @Description Publishes a job message to RabbitMQ with connection parameters
// @Tags Publish
// @Accept json
// @Produce json
// @Param host header string true "RabbitMQ host"
// @Param port header string true "RabbitMQ port"
// @Param user header string true "RabbitMQ user"
// @Param password header string true "RabbitMQ password"
// @Param job body model.Job true "Job object to publish"
// @Success 201 {object} map[string]interface{} "Publish success response"
// @Failure 400 {object} map[string]interface{} "Error response"
// @Router /publish [post]
func RegisterPublishRoutes(r *gin.Engine) {
	r.POST("/publish", controllers.Publish)
}
