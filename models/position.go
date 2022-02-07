package models

import (
	"time"
)

type Position struct {
	ID          	int		`json:"id" gorm:"primaryKey"`
	Latitude		float32		`json:"latitude"`
	Longitude		float32 	`json:"longitude"`
	CurrentSpeed	int			`json:"current_speed"`
	MaxSpeed		int			`json:"max_speed"`
	CreatedAt		time.Time 	`json:"timestamp"`
	Vehicle			int		`json:"vehicle"`
}