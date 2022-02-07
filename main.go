package main

import (
	"vts_api/consumer"
	"vts_api/database"
	"vts_api/server"
)

func main () {
	
	database.StartDatabase()
	go consumer.ConsumerInit()

	server := server.NewServer()
	server.Run()

}