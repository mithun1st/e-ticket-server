package subusermodel

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
