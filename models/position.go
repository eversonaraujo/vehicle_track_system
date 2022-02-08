package models

import (
	"regexp"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Latitude		float32		`json:"latitude" valid:"required,latitude"`
	Longitude		float32 	`json:"longitude" valid:"required,longitude"`
	CurrentSpeed	float32		`json:"current_speed" valid:"required,float"`
	MaxSpeed		float32		`json:"max_speed" valid:"optional,float"`
	Timestamp		string		`json:"timestamp" valid:"required,iso8601date"`
	VehicleID		int			`json:"vehicle_id"`
}

func (position *Position) Validate() error {

	_, err:= govalidator.ValidateStruct(position)
	
	if err != nil {
		return err
	}

	return nil
}

func init () {
	govalidator.TagMap["iso8601date"] = govalidator.Validator(func(str string) bool {
		ISO8601DateRegexString := "^(\\d{4})(-(0[1-9]|1[0-2])(-([12]\\d|0[1-9]|3[01]))([T\\s]((([01]\\d|2[0-3])((:)[0-5]\\d))([\\:]\\d+)?)?(:[0-5]\\d([\\.]\\d+)?)?([zZ]|([\\+-])([01]\\d|2[0-3]):?([0-5]\\d)?)?)?)$"
		ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
		return ISO8601DateRegex.MatchString(str)
	})
}
