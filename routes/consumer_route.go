package routes

import (
	"AgentApiGo/controllers"
	"github.com/gin-gonic/gin"
)

// Consume jobs of RabbitMQ
// @Summary Consume a job
// @Description Publishes a job message to RabbitMQ with connection parameters
// @Tags Consume
// @Accept json
// @Produce json
// @Param host header string true "RabbitMQ host"
// @Param port header string true "RabbitMQ port"
// @Param user header string true "RabbitMQ user"
// @Param password header string true "RabbitMQ password"
// @Param queue query string true "RabbitMQ queue"
// @Success 201 {object} map[string]string "Consumer success"
// @Failure 400 {object} map[string]interface{} "Error response"
// @Router /consumer/start [post]
func RegisterConsumerRoutes(r *gin.Engine) {
	//consumerService := service.NewConsumerService()
	//consumeController := controllers.NewConsumerController(consumerService)
	r.POST("/consumer/start", controllers.Consumer)
}

// {
// 	"server": "172.20.240.1",
// 	"script": "mkdir c:\\temp\\kaio3",
// 	"cmdExecute": "true"
//  }
