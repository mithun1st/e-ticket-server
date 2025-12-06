package routecountermodel

type CounterEntity struct {
	Id             int     `json:"id"`
	FkAssignUserId *int    `json:"fkAssignUserId"`
	Name           string  `json:"name"`
	Address        *string `json:"address"`
	Lat            *string `json:"lat"`
	Long           *string `json:"long"`
	IsActive       bool    `json:"isActive"`
}
