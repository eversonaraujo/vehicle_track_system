package controllers

import (
	"fmt"
	"strconv"
	"vts_api/database"
	"vts_api/models"

	"github.com/gin-gonic/gin"
)

// /fleet/{FLEET_ID}/alerts to see all alerts from this Fleet
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
		c.JSON(400, gin.H { "errors": "Invalid parameter Fleet.ID: " +  err.Error() })
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
		
	var alert models.Alert
	err = c.ShouldBindJSON(&alert)
	alert.FleetID = int(fleet.ID)

	if err != nil {
		c.JSON(400, gin.H { "errors": "Invalid JSON"})
		return
	}

	// Validating
	err = alert.Validate() 
	if err != nil {
		c.JSON(400, gin.H { "errors": err.Error() })
		return
	}

	// Check if this alert (url) already exists as alert
	var count int64
	database.GetDabatase().Table("alerts").Where("webhook = ? AND fleet_id = ?", alert.Webhook, alert.FleetID).Count(&count)
	
	if (count > 0) {
		c.JSON(400, gin.H { "errors": fmt.Sprintf("This alert (\"%s\") already exists", alert.Webhook )})
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