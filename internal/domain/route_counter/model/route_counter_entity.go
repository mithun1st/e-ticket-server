package routecountermodel

type RouteCounterEntity struct {
	CounterId int      `json:"counterId"`
	Duration  *int     `json:"duration"`
	Cost      *float64 `json:"cost"`
	Serial    int      `json:"serial"`
}
