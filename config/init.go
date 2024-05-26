package config

import (
	"astro_drone/app/controller"
	"astro_drone/app/repository"
	"astro_drone/app/service"
)

type Initialization struct {
	DroneRepo repository.DroneRepository
	DroneSvc  service.DroneService
	DroneCtrl controller.DroneController
}

func NewInitialization(droneRepo repository.DroneRepository,
	droneService service.DroneService,
	droneCtrl controller.DroneController) *Initialization {
	return &Initialization{
		DroneRepo: droneRepo,
		DroneSvc:  droneService,
		DroneCtrl: droneCtrl,
	}
}
