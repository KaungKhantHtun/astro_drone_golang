package repository

import (
	"astro_drone/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DroneRepository interface {
	FindAllDrone() ([]dao.Drone, error)
	FindDroneById(id int) (dao.Drone, error)
	Save(drone *dao.Drone) (dao.Drone, error)
	SaveActionLog(actionLog *dao.DroneActionLog) (dao.DroneActionLog, error)
	FindFirstActionLog(id int) (dao.DroneActionLog, error)
	FindLastActionLog(id int) (dao.DroneActionLog, error)
	FindActionLogs(id int) ([]dao.DroneActionLog, error)
}

type DroneRepositoryImpl struct {
	db *gorm.DB
}

func (d DroneRepositoryImpl) Save(drone *dao.Drone) (dao.Drone, error) {
	var err = d.db.Save(drone).Error
	if err != nil {
		log.Error("Got an error when save drone. Error: ", err)
		return dao.Drone{}, err
	}
	return *drone, nil
}

func (d DroneRepositoryImpl) FindAllDrone() ([]dao.Drone, error) {
	var drones []dao.Drone

	var err = d.db.Find(&drones).Error
	if err != nil {
		log.Error("Got an error finding all drone reports. Error: ", err)
		return nil, err
	}

	return drones, nil
}

func (d DroneRepositoryImpl) FindDroneById(id int) (dao.Drone, error) {
	drone := dao.Drone{
		Id: id,
	}
	err := d.db.First(&drone).Error
	if err != nil {
		log.Error("Got and error when find drone by id. Error: ", err)
		return dao.Drone{}, err
	}
	return drone, nil
}

func (d DroneRepositoryImpl) SaveActionLog(actionLog *dao.DroneActionLog) (dao.DroneActionLog, error) {
	var err = d.db.Save(actionLog).Error
	if err != nil {
		log.Error("Got an error when save action log. Error: ", err)
		return dao.DroneActionLog{}, err
	}
	return *actionLog, nil
}

func (d DroneRepositoryImpl) FindActionLogs(id int) ([]dao.DroneActionLog, error) {

	var actionLogs []dao.DroneActionLog

	log.Info(id)

	var err = d.db.Where(dao.DroneActionLog{DroneID: id}).Find(&actionLogs).Error

	if err != nil {
		log.Error("Got an error when find action logs by id. Error: ", err)
		return nil, err
	}
	return actionLogs, nil
}

func (d DroneRepositoryImpl) FindFirstActionLog(id int) (dao.DroneActionLog, error) {

	startActionLog := dao.DroneActionLog{
		DroneID: id,
	}

	var err = d.db.First(&startActionLog).Error

	if err != nil {
		log.Error("Got an error when find first action log by id. Error: ", err)
		return dao.DroneActionLog{}, err
	}
	return startActionLog, nil
}

func (d DroneRepositoryImpl) FindLastActionLog(id int) (dao.DroneActionLog, error) {

	lastActionLog := dao.DroneActionLog{
		DroneID: id,
	}

	var err = d.db.Last(&lastActionLog).Error

	if err != nil {
		log.Error("Got an error when find last action log by id. Error: ", err)
		return dao.DroneActionLog{}, err
	}
	return lastActionLog, nil
}

func DroneRepositoryInit(db *gorm.DB) *DroneRepositoryImpl {
	db.AutoMigrate(&dao.Drone{}, &dao.DroneActionLog{})
	return &DroneRepositoryImpl{
		db: db,
	}
}
