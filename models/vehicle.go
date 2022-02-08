package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Name        string				`json:"name" valid:"required"`
	MaxSpeed	float32				`json:"max_speed" valid:"optional,float"`
	FleetID		int					`json:"fleet_id" valid:"required,int"`
	CreatedAt time.Time 			`json:"created"`
	UpdatedAt time.Time 			`json:"updated"`
 	DeletedAt gorm.DeletedAt 		`gorm:"index" json:"deleted"`
}

func (vehicle *Vehicle) Validate() error {

	_, err:= govalidator.ValidateStruct(vehicle)
	
	if err != nil {
		return err
	}

	return nil
}