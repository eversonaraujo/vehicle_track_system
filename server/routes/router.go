package routes

import (
	"vts_api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes (router *gin.Engine) *gin.Engine {
	main := router
	{	
		main.GET("/", controllers.Hello)
		main.GET("/database", controllers.Reset)

		vehicle := main.Group("vehicles")
		{
			vehicle.GET("/", controllers.VehicleList)
			vehicle.POST("/", controllers.VehicleCreate)

			vehicle.GET("/:id/positions", controllers.PositionList)
			vehicle.POST("/:id/positions", controllers.PositionCreate)
		}
		
		fleet := main.Group("fleets")
		{
			fleet.GET("/", controllers.FleetList)
			fleet.POST("/", controllers.FleetCreate)
			
			fleet.GET("/:id/alerts", controllers.AlertList)
			fleet.POST("/:id/alerts", controllers.AlertCreate)
		}
	}
	
	return router
}