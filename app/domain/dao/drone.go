package dao

type Drone struct {
	Id                        int    `gorm:"column:id; primary_key; not null" json:"id"`
	IsCompleted               bool   `gorm:"column:is_completed; default:false" json:"is_completed" `
	FinalDirectionAndPosition string `gorm:"column:final_direction_and_position" json:"final_direction_and_position"`
	StartAndFinishedTime      string `gorm:"column:start_and_finished_time" json:"start_and_finished_time"`
	BaseModel
}
