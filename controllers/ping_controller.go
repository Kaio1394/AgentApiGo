package controllers

import (
	"AgentApiGo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	s *service.PingService
}

func NewPingController(s *service.PingService) *PingController {
	return &PingController{s: s}
}

func (ps *PingController) Ping(c *gin.Context) {
	ip, host, os := ps.s.GetInformationMachine()
	c.JSON(http.StatusOK, gin.H{
		"host": ip,
		"IP":   host,
		"OS":   os,
	})
}
