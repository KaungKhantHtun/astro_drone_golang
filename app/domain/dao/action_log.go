package dao

type DroneActionLog struct {
	ID                   int    `gorm:"column:id; primary_key; not null" json:"id"`
	DroneID              int    `gorm:"column: drone_id; not null" json:"drone_id"`
	Action               string `gorm:"column: action" json:"action"`
	Times                int64  `gorm:"column: times" json:"times"`
	DirectionAndPosition string `gorm:"column:direction_and_position" json:"direction_and_position"`
	Timestampt           int    `gorm:"column:timestampt; autoCreateTime:milli" json:"timestampt"`
	BaseModel
}
