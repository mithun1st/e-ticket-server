package routecountermodel

type RouteCounterUri struct {
	CompanyId int `uri:"companyId" binding:"required"`
}

// type RouteCounterQuery struct {
// 	Q1 string `form:"q1" binding:"omitempty,max=4"`
// 	Q2 int    `form:"q2" binding:"omitempty,lte=9"`
// }

// type RouteCounterRequest struct {
// 	Id int    `json:"id" binding:"omitempty,gte=18,lte=21"`
// 	B1 string `json:"b1" binding:"required,email"`
// 	B2 string `json:"b2" binding:"required,min=8,max=12"`
// 	B4 string `json:"b4" binding:"omitempty,len=4"`
// 	B5 bool   `json:"b5" binding:"omitempty"`
// }
