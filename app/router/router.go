package router

import (
	"astro_drone/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		drone := api.Group("/drone")

		drone.POST("", init.DroneCtrl.AddDroneData)
		drone.GET("", init.DroneCtrl.GetAllDroneData)
		drone.GET("/:id", init.DroneCtrl.GetDroneById)
	}

	return router
}
