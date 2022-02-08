package controllers

import (
	"vts_api/database"

	"github.com/gin-gonic/gin"
)

func Hello (c *gin.Context) {
	c.JSON(200, "Welcome to Vehicle Track System")
}

func TestPost (c *gin.Context) {
	c.JSON(200, "")
}

func Reset (c *gin.Context) {

	database.Reset()
	c.JSON(200, "Database Droped")
}