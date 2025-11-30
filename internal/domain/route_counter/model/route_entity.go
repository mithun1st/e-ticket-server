package routecountermodel

type RouteEntity struct {
	Id            int     `json:"id"`
	Fk_company_id int     `json:"fk_company_id"`
	Name          string  `json:"name"`
	Note          *string `json:"note"`
	Is_active     bool    `json:"is_active"`
}
