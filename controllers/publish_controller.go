package controllers

import (
	"AgentApiGo/helper"
	"AgentApiGo/logger"
	"AgentApiGo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {

	var job model.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	host := c.DefaultQuery("host", "")
	portStr := c.DefaultQuery("port", "")
	user := c.DefaultQuery("user", "")
	password := c.DefaultQuery("password", "")

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

	var rabbit_config helper.IRabbit

	rabbit_config = helper.Rabbit{
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

	rabbit_config.SendMessage(job, "Job.Schedule.Test", con)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Publish Job.",
		"job":     job,
	})
	logger.Log.Info("Sended message to queue Job.Schedule.Test")
}
