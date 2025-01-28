package model

type Job struct {
	Name        string `json:"name" binding:"required"`
	Server      string `json:"server" binding:"required"`
	Description string `json:"description" binding:"required"`
	CmdExecute  string `json:"cmdExecute" binding:"required"`
	Script      string `json:"script" binding:"required"`
	Date        string `json:"date" binding:"required"`
}
