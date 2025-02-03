package service

import (
	"AgentApiGo/helper"
	"AgentApiGo/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConsumerService struct {
	rabbit *helper.Rabbit
}

func NewConsumerService(rabbit *helper.Rabbit) *ConsumerService {
	return &ConsumerService{rabbit: rabbit}
}

func (cs *ConsumerService) Consumer(c *gin.Context, queue string) {
	con, err := cs.rabbit.Connection()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Connection error.",
			"ParamsConnection": cs.rabbit,
		})
		logger.Log.Error("Connection error.")
		return
	}

	go func() {
		cs.rabbit.Consumer(queue, con)
	}()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Consumer started.",
	})
}
