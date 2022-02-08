package controllers

import (
	"strconv"
	"vts_api/database"
	"vts_api/models"

	"github.com/gin-gonic/gin"
)

func VehicleList (c *gin.Context) {
	
	vehicles := []models.Vehicle{}
	var err = database.GetDabatase().Find(&vehicles).Error

	if err != nil {
		c.JSON(400, gin.H { "errors": "Failed on loading vehicles"})
		return
	}
	
	c.JSON(200, vehicles)
}

func VehicleCreate (c *gin.Context) {
		
	// Don`t forget the validation
	var vehicle models.Vehicle
	err := c.ShouldBindJSON(&vehicle)

	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalid JSON"})
		return
	}
	
	// Validating
	err = vehicle.Validate() 
	if err != nil {
		c.JSON(400, gin.H { "errors": err.Error() })
		return
	}

	// Check if Freet exists
	var fleet models.Fleet
	err = database.GetDabatase().First(&fleet, vehicle.FleetID).Error
	if err != nil {
		c.JSON(400, gin.H {
			"errors": "Not found Fleet by ID: " +  strconv.FormatInt(int64(vehicle.FleetID), 10),
		})
		
		return
	}	

	err = database.GetDabatase().Create(&vehicle).Error

	if err != nil {

		c.JSON(400, gin.H {
			 "errors": "Cannot create a vehicle, error: "+ err.Error(),
		})
		return
	}

	c.JSON(201, gin.H { "id" : vehicle.ID })
}