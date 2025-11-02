package subusermodel

// type SubUserHeader struct {
// 	Auth string `header:"Authorization"`
// }

// type SubUserPath struct {
// 	P1 string `uri:"subUserId" binding:"omitempty,len=4"`
// }

// type SubUserQuery struct {
// 	Q1 string `form:"q1" binding:"omitempty,max=4"`
// 	Q2 int    `form:"q2" binding:"omitempty,lte=9"`
// }

// type SubUserRequest struct {
// 	Id int    `json:"id" binding:"omitempty,gte=18,lte=21"`
// 	B1 string `json:"b1" binding:"required,email"`
// 	B2 string `json:"b2" binding:"required,min=8,max=12"`
// 	B4 string `json:"b4" binding:"omitempty,len=4"`
// 	B5 bool   `json:"b5" binding:"omitempty"`
// }

type SubUserQuery struct {
	CompanyId int  `form:"companyId" binding:"required"`
	Role      *int `form:"role" binding:"omitempty"`
}

type SubUserCreateRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  *string `json:"lastName" binding:"omitempty"`
	Phone     string  `json:"phone" binding:"required"`
	Email     *string `json:"email" binding:"omitempty"`

	CompanyId int `json:"companyId" binding:"required"`
	Role      int `json:"role" binding:"required"`
}
