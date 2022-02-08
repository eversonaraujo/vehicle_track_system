package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)


type Fleet struct {
	gorm.Model
	Name		string		`json:"name" valid:"required"`
	MaxSpeed	float32		`json:"max_speed" valid:"required,range(1|350),float"`
}

func (fleet *Fleet) Validate() error {

	_, err:= govalidator.ValidateStruct(fleet)
	
	if err != nil {
		return err
	}

	return nil
}