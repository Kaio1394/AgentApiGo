package main

import (
	"AgentApiGo/routes"

	"github.com/gin-gonic/gin"

	// pacote correto para os arquivos Swagger
	_ "AgentApiGo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // para integrar o Swagger ao Gin
)

// @title AgentApiGo
// @version 1.0
// @description This is a sample API to demonstrate Swagger with Gin.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterPingRoutes(r)
	routes.RegisterPublishRoutes(r)

	r.Run(":8080")
}
