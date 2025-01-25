package controllers

import (
	"AgentApiGo/logger"
	"AgentApiGo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Consumer(c *gin.Context) {
	host := c.DefaultQuery("host", "")
	portStr := c.DefaultQuery("port", "")
	user := c.DefaultQuery("user", "")
	password := c.DefaultQuery("password", "")
	queue := c.DefaultQuery("queue", "")

	if host == "" || portStr == "" || user == "" || password == "" {
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

	var rabbit_config model.IRabbit

	rabbit_config = model.Rabbit{
		Host:     host,
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
