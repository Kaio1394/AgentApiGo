package controllers

import (
	"AgentApiGo/helper"
	"AgentApiGo/logger"
	"AgentApiGo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConsumerController struct {
	service *service.ConsumerService
}

func NewConsumerController(service *service.ConsumerService) *ConsumerController {
	return &ConsumerController{}
}

func Consumer(c *gin.Context) {
	server := c.GetHeader("server")
	portStr := c.GetHeader("port")
	user := c.GetHeader("user")
	password := c.GetHeader("password")
	queue := c.DefaultQuery("queue", "")

	if server == "" || portStr == "" || user == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required parameters (host, user, password, port).",
		})
		logger.Log.Error("Missing required parameters (host, user, password, port).")
		return
	}

	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid port parameter.",
		})
		logger.Log.Error("Invalid port parameter.")
		return
	}

	var rabbit_config helper.IRabbit

	rabbit_config = helper.Rabbit{
		Host:     server,
		Port:     uint32(port),
		User:     user,
		Password: password,
	}

	if rabbit_config.HasEmptyParams() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required parameters (host, user, password, port).",
		})
		logger.Log.Error("Missing required parameters (host, user, password, port).")
		return
	}

	con, err := rabbit_config.Connection()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Connection error.",
			"ParamsConnection": rabbit_config,
		})
		logger.Log.Error("Connection error.")
		return
	}

	go func() {
		rabbit_config.Consumer(queue, con)
	}()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Consumer started.",
	})
}
