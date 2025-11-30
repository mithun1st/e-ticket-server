package routecountermodel

type CounterEntity struct {
	Id             int     `json:"id"`
	FkCompanyId    int     `json:"fk_company_id"`
	FkAssignUserId *int    `json:"fk_assign_user_id"`
	Name           string  `json:"name"`
	Address        *string `json:"address"`
	Lat            *string `json:"lat"`
	Long           *string `json:"long"`
	IsActive       bool    `json:"is_active"`
}
