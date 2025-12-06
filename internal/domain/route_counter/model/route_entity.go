package routecountermodel

type RouteEntity struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Note     *string `json:"note"`
	IsActive bool    `json:"isActive"`
}
