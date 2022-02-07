package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vts_api/models"

	"github.com/gin-gonic/gin"
)

type NotifyBody struct {
	Webhook string 				`json:"webhook"`
	Position models.Position	`json:"position"`
}

// Return true on success otherwise false
func Notify (message []byte) bool {
	
	var messageJson = NotifyBody{}

	err := json.Unmarshal(message, &messageJson)
	
	if err != nil {
		log.Fatal("Failed on decoding message #1")
		return false
	}

	messageBody, err := json.Marshal(
		gin.H { 
			"id"			: messageJson.Position.ID,
			"vehicle_id"	: messageJson.Position.VehicleID,  
			"timestamp" 	: messageJson.Position.Timestamp,
			"latitude" 		: messageJson.Position.Latitude,
			"longitude" 	: messageJson.Position.Longitude,
			"current_speed" : messageJson.Position.CurrentSpeed,
			"max_speed" 	: messageJson.Position.MaxSpeed,		
	})
	
	if err != nil {
		log.Fatal("Failed on decoding message #2")
		return false
	}

	resp, err := http.Post(
		messageJson.Webhook, 
		"application/json", 
		bytes.NewBuffer(messageBody),
	)

	if err != nil {
		fmt.Println("Failed: " + err.Error())
		return false
	}
	
	if resp.StatusCode == 200 {
		fmt.Println("Mensage received!")
		return true
	} 
		
	fmt.Println("Ops, failed on sending message")
	return false
}