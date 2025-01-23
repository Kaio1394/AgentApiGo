package routes

import (
	"AgentApiGo/controllers"

	"github.com/gin-gonic/gin"
)

// @Summary Publicar algo
// @Description Endpoint para publicar dados
// @Tags Publish
// @Accept json
// @Produce json
// @Param request body map[string]interface{} true "Corpo da requisição"
// @Success 201 {object} map[string]interface{}
// @Router /publish [post]
func RegisterPublishRoutes(r *gin.Engine) {
	r.POST("/publish", controllers.Publish)
}
