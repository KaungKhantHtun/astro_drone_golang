package service

import (
	"astro_drone/app/constant"
	"astro_drone/app/domain/dao"
	"astro_drone/app/domain/dto"
	"astro_drone/app/pkg"
	"astro_drone/app/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type DroneService interface {
	GetAllDrone(c *gin.Context)
	GetDroneById(c *gin.Context)
	AddDroneData(c *gin.Context)
}

type DroneServiceImpl struct {
	droneRepository repository.DroneRepository
}

func (d DroneServiceImpl) AddDroneData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data drone")
	var request []dto.AddDroneRequestDTO
	var drone dao.Drone
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	var result []dao.Drone
	for _, req := range request {
		drone = dao.Drone{Id: req.DroneID}
		data, err := d.droneRepository.Save(&drone)
		if err != nil {
			log.Error("Happened error when saving data to database. Error", err)
			pkg.PanicException(constant.UnknownError)
		}
		result = append(result, data)
	}

	CreateInstructionLog(request, d)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, result))
}

func (d DroneServiceImpl) GetDroneById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get drone by id")
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := d.droneRepository.FindDroneById(id)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	actionLogData, err := d.droneRepository.FindActionLogs(id)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	droneDetailResponse := dto.DroneDetailResponse{
		Id:                        data.Id,
		IsCompleted:               data.IsCompleted,
		FinalDirectionAndPosition: data.FinalDirectionAndPosition,
		StartAndFinishedTime:      data.StartAndFinishedTime,
		ActionLogs:                actionLogData,
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, droneDetailResponse))
}

func (d DroneServiceImpl) GetAllDrone(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data drone")

	data, err := d.droneRepository.FindAllDrone()
	if err != nil {
		log.Error("Happened Error when find all drone data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func DroneServiceInit(droneRepository repository.DroneRepository) *DroneServiceImpl {
	return &DroneServiceImpl{
		droneRepository: droneRepository,
	}
}

func Move(current dao.Position, action string, times int) dao.Position {
	var currentPosition dao.Position = current
	var currentIndex int
	var directions = []string{"N", "E", "S", "W"}

	for i, dir := range directions {
		if dir == currentPosition.Direction {
			currentIndex = i
			break
		}
	}

	var newIndex int
	if action == "L" {
		newIndex = (currentIndex - times + len(directions)) % len(directions)
		currentPosition.Direction = directions[newIndex]
	} else if action == "R" {
		newIndex = (currentIndex + times) % len(directions)
		currentPosition.Direction = directions[newIndex]
	} else if action == "F" {
		switch currentPosition.Direction {
		case "N":
			currentPosition.Y += times
		case "E":
			currentPosition.X += times
		case "S":
			if currentPosition.Y >= times {
				currentPosition.Y -= times
			} else {
				currentPosition.Y = 0
			}
		case "W":
			if currentPosition.X >= times {
				currentPosition.X -= times
			} else {
				currentPosition.X = 0
			}
		}
	}
	return currentPosition
}

func CreateInstructionLog(drones []dto.AddDroneRequestDTO, d DroneServiceImpl) {
	for _, drone := range drones {
		var currentPosition dao.Position
		currentPosition.X = 0
		currentPosition.X = 0
		currentPosition.Direction = "N"
		initialActionLog := dao.DroneActionLog{DroneID: drone.DroneID}
		d.droneRepository.SaveActionLog(&initialActionLog)

		for i, instruction := range drone.Instructions {
			currentPosition = Move(currentPosition, instruction.Action, int(instruction.Times))
			str := currentPosition.String()
			currentInfo := dao.DroneActionLog{DroneID: drone.DroneID, Action: instruction.Action, DirectionAndPosition: str, Times: instruction.Times}
			d.droneRepository.SaveActionLog(&currentInfo)

			if i == len(drone.Instructions)-1 {
				var firstInfo dao.DroneActionLog
				var lastInfo dao.DroneActionLog
				updateDrone := dao.Drone{
					Id: drone.DroneID,
				}

				firstInfo, firstInfoErr := d.droneRepository.FindFirstActionLog(drone.DroneID)
				if firstInfoErr != nil {
					log.Error("Happened Error when find first action log by id. Error: ", firstInfoErr)
					pkg.PanicException(constant.UnknownError)
				}

				lastInfo, endInfoErr := d.droneRepository.FindLastActionLog(drone.DroneID)
				if endInfoErr != nil {
					log.Error("Happened Error when find last action log by id. Error: ", endInfoErr)
					pkg.PanicException(constant.UnknownError)
				}

				updateDrone.IsCompleted = true
				updateDrone.FinalDirectionAndPosition = str
				updateDrone.StartAndFinishedTime = fmt.Sprintf("%v - %v", firstInfo.Timestampt, lastInfo.Timestampt)
				d.droneRepository.Save(&updateDrone)
				log.Info(updateDrone)

			}

		}
	}
}
