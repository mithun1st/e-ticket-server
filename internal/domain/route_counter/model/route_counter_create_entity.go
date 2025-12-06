package routecountermodel

type RouteCounterCreateEntity struct {
	CounterId int      `json:"counterId" binding:"required"`
	Duration  *int     `json:"duration" binding:"omitempty"`
	Cost      *float64 `json:"cost" binding:"omitempty"`
	Serial    int      `json:"serial" binding:"required"`
}
