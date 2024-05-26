package controller

import (
	"astro_drone/app/service"

	"github.com/gin-gonic/gin"
)

type DroneController interface {
	AddDroneData(c *gin.Context)
	GetAllDroneData(c *gin.Context)
	GetDroneById(c *gin.Context)
}

type DroneControllerImpl struct {
	svc service.DroneService
}

func (d DroneControllerImpl) AddDroneData(c *gin.Context) {
	d.svc.AddDroneData(c)
}

func (d DroneControllerImpl) GetAllDroneData(c *gin.Context) {
	d.svc.GetAllDrone(c)
}

func (d DroneControllerImpl) GetDroneById(c *gin.Context) {
	d.svc.GetDroneById(c)
}

func DroneControllerInit(droneService service.DroneService) *DroneControllerImpl {
	return &DroneControllerImpl{
		svc: droneService,
	}
}
