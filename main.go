package main

import (
	"AgentApiGo/logger"
	"AgentApiGo/routes"

	"github.com/gin-gonic/gin"

	// pacote correto para os arquivos Swagger
	_ "AgentApiGo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // para integrar o Swagger ao Gin
)

var port string = "8080"

// @title AgentApiGo
// @version 1.0
// @description This is a sample API to demonstrate Swagger with Gin.
// @host localhost:8080
// @BasePath /
func main() {
	logger.Log.Info("Start application...")
	r := gin.Default()

	logger.Log.Info("Calling swagger endpoint.")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterPingRoutes(r)
	routes.RegisterPublishRoutes(r)
	routes.RegisterConsumerRoutes(r)

	logger.Log.Info("Project listening in port " + port)
	r.Run(":" + port)
}
