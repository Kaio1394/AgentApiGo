package model

type Job struct {
	Id          int    `json:"id" gorm:"column:id" example:"1"`
	Name        string `json:"name" gorm:"column:name" example:"Backup Job"`
	Description string `json:"description" gorm:"column:description,omitempty" example:"Daily backup job"`
	CmdExecute  bool   `json:"cmdExecute" gorm:"column:cmd_execute" example:"true"`
	Script      string `json:"script" gorm:"column:script,omitempty" example:"backup.sh"`
	Date        string `json:"date" gorm:"column:date,omitempty" example:"2025-02-04T15:04:05Z"`
	Hold        bool   `json:"hold" gorm:"column:hold" example:"false"`
	Priority    int    `json:"priority" gorm:"column:priority" example:"1"`
	ServerId    int    `json:"serverId" gorm:"column:server_id" example:"2"`
}
