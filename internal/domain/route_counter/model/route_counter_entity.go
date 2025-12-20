package routecountermodel

type RouteCounterEntity struct {
	CounterId   int      `json:"counterId"`
	CounterName string   `json:"counterName"`
	Distance    *float64 `json:"distance"`
	Duration    *int     `json:"duration"`
	Serial      int      `json:"serial"`
}
