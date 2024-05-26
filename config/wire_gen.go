// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"astro_drone/app/controller"
	"astro_drone/app/repository"
	"astro_drone/app/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	droneRepositoryImpl := repository.DroneRepositoryInit(gormDB)
	droneServiceImpl := service.DroneServiceInit(droneRepositoryImpl)
	droneControllerImpl := controller.DroneControllerInit(droneServiceImpl)
	initialization := NewInitialization(droneRepositoryImpl, droneServiceImpl, droneControllerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var droneServiceSet = wire.NewSet(service.DroneServiceInit, wire.Bind(new(service.DroneService), new(*service.DroneServiceImpl)))

var droneRepoSet = wire.NewSet(repository.DroneRepositoryInit, wire.Bind(new(repository.DroneRepository), new(*repository.DroneRepositoryImpl)))

var droneCtrlSet = wire.NewSet(controller.DroneControllerInit, wire.Bind(new(controller.DroneController), new(*controller.DroneControllerImpl)))