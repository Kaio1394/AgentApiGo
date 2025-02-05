package model

import "time"

type HistoryExecution struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Status   string    `json:"tag"`
	JobId    int       `json:"jobId" gorm:"column:job_id"`
	ServerId int       `json:"serverId" gorm:"column:server_id"`
	CreateAt time.Time `json:"createAt" gorm:"column:create_at;default:CURRENT_TIMESTAMP"`
}

func (HistoryExecution) TableName() string {
	return "h_job_history_execution"
}
