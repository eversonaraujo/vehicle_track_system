package models

type Alert struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	FleetID   	int	`json:"fleet"`
	Webhook 	string  `json:"webhook"`
}