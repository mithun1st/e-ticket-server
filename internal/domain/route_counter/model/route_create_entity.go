package routecountermodel

type RouteCreateEntity struct {
	Name string  `json:"name" binding:"required"`
	Note *string `json:"note" binding:"omitempty"`
}
