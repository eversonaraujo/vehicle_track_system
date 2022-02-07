package models

import "gorm.io/gorm"

type Fleet struct {
	gorm.Model
	Name		string		`json:"name"`
	MaxSpeed	float32		`json:"max_speed"`
}