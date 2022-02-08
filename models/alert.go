package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Alert struct {
	gorm.Model
	Webhook 	string  `json:"webhook" valid:"required,url"`
	FleetID		int 	`json:"fleet_id"`
	// Fleet		Fleet	`json:"fleet"`
}

func (alert *Alert) Validate() error {

	_, err:= govalidator.ValidateStruct(alert)
	
	if err != nil {
		return err
	}

	return nil
}