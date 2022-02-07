package server

import (
	"log"
	"vts_api/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer () Server {
	return Server {
		port: "8000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {

	router := routes.ConfigRoutes(s.server)

	log.Printf("---------------------- READY ----------------------")
	log.Printf("Server is running on %v", s.port)
	log.Fatal(router.Run(":" + s.port ))
}