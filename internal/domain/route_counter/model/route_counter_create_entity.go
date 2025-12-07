package routecountermodel

type RouteCounterCreateEntity struct {
	CounterId int      `json:"counterId" binding:"required"`
	Distance  *float64 `json:"distance" binding:"omitempty"`
	Duration  *int     `json:"duration" binding:"omitempty"`
	Serial    int      `json:"serial" binding:"required"`
}
