package main

import (
	"vts_api/database"
	"vts_api/server"
)

func main () {
	
	database.StartDatabase()
	server := server.NewServer()
	server.Run()

	// From Docs
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()
}