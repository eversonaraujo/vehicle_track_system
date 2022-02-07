package controllers

import (
	"strconv"
	"vts_api/database"
	"vts_api/models"

	"github.com/gin-gonic/gin"
)

func AlertList (c *gin.Context) {
	
	alerts := []models.Alert{}
	var err = database.GetDabatase().Find(&alerts).Error

	if err != nil {
		c.JSON(400, gin.H { "errors": "Failed on loading fleets"})
		return
	}
	
	c.JSON(200, alerts)
}

func AlertCreate (c *gin.Context) {

	fleetID, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalida parameter Fleet.ID: " +  err.Error() })
		return
	}
	
	var fleet models.Fleet
	// Check if Freet exists
	err = database.GetDabatase().First(&fleet, fleetID).Error
	if err != nil {
		c.JSON(400, gin.H {
			"errors": "Not found Fleet by ID: " + c.Param("id"),
		})
		
		return
	}	
		
	// Don`t forget the validation
	var alert models.Alert
	err = c.ShouldBindJSON(&alert)
	alert.FleetID = fleetID

	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalid JSON"})
		return
	}

	err = database.GetDabatase().Create(&alert).Error

	if err != nil {

		c.JSON(400, gin.H {
			 "errors": "Cannot create a alert, error: "+ err.Error(),
		})
		return
	}

	c.JSON(201, gin.H { "id" : alert.ID })
}