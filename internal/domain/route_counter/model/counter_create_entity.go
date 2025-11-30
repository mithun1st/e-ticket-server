package routecountermodel

type CounterCreateEntity struct {
	FkAssignUserId *int    `json:"fk_assign_user_id" binding:"omitempty"`
	Name           string  `json:"name" binding:"required"`
	Address        *string `json:"address" binding:"omitempty"`
	Lat            *string `json:"lat" binding:"omitempty"`
	Long           *string `json:"long" binding:"omitempty"`
}
