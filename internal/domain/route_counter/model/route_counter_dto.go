package routecountermodel

type RouteAndCounterUri struct {
	CompanyId int `uri:"companyId" binding:"required"`
}

type RouteCounterUri struct {
	CompanyId int `uri:"companyId" binding:"required"`
	RouteId   int `uri:"routeId" binding:"required"`
}
