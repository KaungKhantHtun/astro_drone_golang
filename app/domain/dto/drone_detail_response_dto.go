package dto

type DroneDetailResponse struct {
	Id                        int    `json:"id"`
	IsCompleted               bool   `json:"is_completed" `
	FinalDirectionAndPosition string `json:"final_direction_and_position"`
	StartAndFinishedTime      string `json:"start_and_finished_time"`
	ActionLogs                any    `json:"action_logs"`
}
