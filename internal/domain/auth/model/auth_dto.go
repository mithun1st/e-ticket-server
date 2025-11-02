package authmodel

// type AuthHeader struct {
// 	Auth string `header:"Authorization"`
// }

// type AuthPath struct {
// 	P1 string `uri:"authId" binding:"omitempty,len=4"`
// }

// type AuthQuery struct {
// 	Q1 string `form:"q1" binding:"omitempty,max=4"`
// 	Q2 int    `form:"q2" binding:"omitempty,lte=9"`
// }

type AuthQuery struct {
	CompanyId int `form:"companyId" binding:"required"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}
