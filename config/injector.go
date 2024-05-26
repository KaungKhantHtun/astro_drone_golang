//go:build wireinject
// +build wireinject

package config

import (
	"astro_drone/app/controller"
	"astro_drone/app/repository"
	"astro_drone/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var droneServiceSet = wire.NewSet(service.DroneServiceInit,
	wire.Bind(new(service.DroneService), new(*service.DroneServiceImpl)),
)

var droneRepoSet = wire.NewSet(repository.DroneRepositoryInit,
	wire.Bind(new(repository.DroneRepository), new(*repository.DroneRepositoryImpl)),
)

var droneCtrlSet = wire.NewSet(controller.DroneControllerInit,
	wire.Bind(new(controller.DroneController), new(*controller.DroneControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, droneCtrlSet, droneServiceSet, droneRepoSet)
	return nil
}
