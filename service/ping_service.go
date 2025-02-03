package service

import "AgentApiGo/helper"

type PingService struct {
	helper *helper.Helper
}

func NewPingService(helper *helper.Helper) *PingService {
	return &PingService{helper: helper}
}

func (ps *PingService) GetInformationMachine() (string, string, string) {
	return ps.helper.GetIp(), ps.helper.GetHost(), ps.helper.GetOperationSystem()
}
