package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string	`json:"name"`
	MaxSpeed	float32	`json:"max_speed"`
	FleetID		int	`json:"fleet_id"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}