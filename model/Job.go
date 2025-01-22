package model

import "time"

type Job struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Priority    uint32    `json:"priority" binding:"required"`
}
