package main

import (
	"AgentApiGo/routes"

	"github.com/gin-gonic/gin"
)

// @title My API
// @version 1.0
// @description This is a sample API to demonstrate Swagger with Gin.
func main() {
	r := gin.Default()

	routes.RegisterPingRoutes(r)
	routes.RegisterPublishRoutes(r)

	r.Run(":8080")
}
