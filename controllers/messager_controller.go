package controllers

import (
	"vts_api/database"
	"vts_api/models"

	"github.com/gin-gonic/gin"
)

func MessagList (c *gin.Context) {
	
	vehicles := []models.Vehicle{}
	var err = database.GetDabatase().Find(&vehicles).Error

	if err != nil {
		c.JSON(400, gin.H { "errors": "Failed on loading vehicles"})
		return
	}
	
	c.JSON(200, vehicles)
}