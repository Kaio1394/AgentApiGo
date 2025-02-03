package routes

import (
	"AgentApiGo/controllers"
	"AgentApiGo/helper"
	"AgentApiGo/service"

	"github.com/gin-gonic/gin"
)

// @Summary Ping endpoint
// @Description Verifica se a API está ativa
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func RegisterPingRoutes(r *gin.Engine) {
	pingService := service.NewPingService(&helper.Helper{})
	pingController := controllers.NewPingController(pingService)
	r.GET("/ping", pingController.Ping)
}
