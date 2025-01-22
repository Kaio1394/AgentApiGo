package controllers

import (
	"AgentApiGo/model"
	"log"
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
		return
	}

	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid port parameter.",
		})
		return
	}

	rabbit_config := model.Rabbit{
		Host:     host,
		Port:     uint32(port),
		User:     user,
		Password: password,
	}

	if rabbit_config.HasEmptyParams() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required parameters (host, user, password, port).",
		})
		return
	}

	con, err := rabbit_config.Connection()
	if err != nil {
		log.Fatal("Connection error!")
	}

	rabbit_config.SendMessage(&job, "Job.Schedule.Test", con)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Publish Job.",
		"job":     job,
	})
}
