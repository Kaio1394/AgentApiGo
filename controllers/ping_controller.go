package controllers

import (
	"AgentApiGo/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"host": helper.GetHost(),
		"IP":   helper.GetIp(),
		"OS":   helper.GetOperationSystem(),
	})
}
