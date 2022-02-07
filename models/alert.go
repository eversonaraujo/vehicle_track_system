package models

import "gorm.io/gorm"

type Alert struct {
	gorm.Model
	Webhook 	string  `json:"webhook"`
	FleetID		int 	`json:"fleet_id"`
	Fleet		Fleet	`json:"fleet"`
}