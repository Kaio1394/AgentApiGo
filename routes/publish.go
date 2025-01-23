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
// @Param host query string true "RabbitMQ host"
// @Param port query string true "RabbitMQ port"
// @Param user query string true "RabbitMQ user"
// @Param password query string true "RabbitMQ password"
// @Param job body model.Job true "Job object to publish"
// @Success 201 {object} map[string]interface{} "Publish success response"
// @Failure 400 {object} map[string]interface{} "Error response"
// @Router /publish [post]
func RegisterPublishRoutes(r *gin.Engine) {
	r.POST("/publish", controllers.Publish)
}
