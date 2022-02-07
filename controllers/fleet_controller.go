package controllers

import (
	"vts_api/database"
	"vts_api/models"

	"github.com/gin-gonic/gin"
)

func FleetList (c *gin.Context) {
	
	fleets := []models.Fleet{}
	var err = database.GetDabatase().Find(&fleets).Error

	if err != nil {
		c.JSON(400, gin.H { "errors": "Failed on loading fleets"})
		return
	}
	
	c.JSON(200, fleets)
}

func FleetCreate (c *gin.Context) {

	// Don`t forget the validation
	var fleet models.Fleet
	err := c.ShouldBindJSON(&fleet) 
	
	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalid JSON"})
		return
	}

	err = database.GetDabatase().Create(&fleet).Error

	if err != nil {

		c.JSON(400, gin.H {
			 "errors": "Cannot create a fleet, error: "+ err.Error(),
		})
		
		return
	}

	c.JSON(200, gin.H { "id" : fleet.ID })
}