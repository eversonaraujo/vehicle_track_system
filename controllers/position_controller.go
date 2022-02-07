package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"vts_api/database"
	"vts_api/messager"
	"vts_api/models"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func PositionList (c *gin.Context) {
	
	positions := []models.Position{}
	var err = database.GetDabatase().Find(&positions).Error

	if err != nil {
		c.JSON(400, gin.H { "errors": "Failed on loading positions"})
		return
	}
	
	c.JSON(200, positions)
}

func PositionCreate (c *gin.Context) {

	vehicleID, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalid parameter Vehicle.ID: " +  err.Error() })
		return
	}

	// Don`t forget the validation
	var position models.Position
	err = c.ShouldBindJSON(&position)

	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalid JSON"})
		return
	}
	
	// Check if Freet exists
	var vehicle models.Vehicle
	err = database.GetDabatase().First(&vehicle, vehicleID).Error
	
	if err != nil {
		c.JSON(400, gin.H { "errors": "Vehicle not exists" })
		return
	}

	position.VehicleID = vehicleID
	
	if position.MaxSpeed == 0 && vehicle.MaxSpeed == 0  {
		
		var fleet models.Fleet
		err = database.GetDabatase().First(&fleet, vehicle.FleetID).Error
		
		if err != nil {
			c.JSON(400, gin.H { "errors": "FleetID not exists" })
			return
		}

		position.MaxSpeed = fleet.MaxSpeed
		return
	} else {
		position.MaxSpeed = vehicle.MaxSpeed
	}

	// Save the VehiclePosition on Database
	err = database.GetDabatase().Create(&position).Error

	if err != nil {

		c.JSON(400, gin.H {
			 "errors": "Cannot create a position, error: "+ err.Error(),
		})
		return
	}
	
	PositionCheckSpeed(position)
	c.JSON(201, gin.H { "id" : position.ID })
}

// Após salvar a posição do veículo, você deverá verificar se a velocidade do veículo é maior do que a
// cadastrada (devemos verificar a velocidade cadastrada no veículo e se não tiver, usar a da frota).
// Caso a velocidade seja maior, deverá enviar as seguintes informações para todos os webhooks
// cadastrados da frota:
// ● POST {url cadastrada}
func PositionCheckSpeed (position models.Position) {
	
	// Speed wrong, too fast
	if position.CurrentSpeed > position.MaxSpeed {
		
		var result []string
		
		database.GetDabatase().Raw("SELECT alerts.webhook FROM alerts JOIN fleets f " +
			" ON f.id = alerts.fleet_id JOIN vehicles v ON v.fleet_id = f.id AND v.id = ?", position.VehicleID).Scan(&result)

		log.Print(strings.Join(result, ", "))
		
		for index := range result {
			
			log.Printf("Publishing message %v...", index)
			
			js, err := json.Marshal(
				gin.H { 
					"webhook"	: result[index],
					"position"	: gin.H {
						"id"			: position.ID,
						"vehicle_id"	: position.VehicleID,  
						"timestamp" 	: position.Timestamp,
						"latitude" 		: position.Latitude,
						"longitude" 	: position.Longitude,
						"current_speed" : position.CurrentSpeed,
						"max_speed" 	: position.MaxSpeed,
					},
			})
	
			if err != nil {
				log.Fatal("Failed on sending message, position to JSON error")
				return
			}

			message := amqp.Publishing {
				ContentType: "application/json",
				Body: []byte(js),
			}

			messager.Publish(message)
		}
			
	}
}