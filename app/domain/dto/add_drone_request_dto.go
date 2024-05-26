package dto

type AddDroneRequestDTO struct {
	DroneID      int           `json:"drone_id"`
	Instructions []Instruction `json:"instructions"`
}

type Instruction struct {
	Action string `json:"action"`
	Times  int64  `json:"times"`
}
