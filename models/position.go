package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Latitude		float32		`json:"latitude"`
	Longitude		float32 	`json:"longitude"`
	CurrentSpeed	float32		`json:"current_speed"`
	MaxSpeed		float32		`json:"max_speed"`
	Timestamp		string		`json:"timestamp"`
	VehicleID		int			`json:"vehicle_id"`
	// Vehicle			Vehicle	
}

// func (position *Position ) AfterCreate(tx *gorm.DB) (err error) {
	
// 	if position.CurrentSpeed > position.MaxSpeed {
// 	// || ( position.CurrentSpeed > position.Vehicle.MaxSpeed) {

		
// 		message := amqp.Publishing {
// 			ContentType: "application/json",
// 			Body: []byte(gin.H{position}),
// 		}

// 		Publisher(message)
// 	}

// 	return
// }