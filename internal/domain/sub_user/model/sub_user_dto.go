package subusermodel

type SubUserUri struct {
	CompanyId int `uri:"companyId" binding:"required"`
}
type SubUserQuery struct {
	Role *int `form:"role" binding:"omitempty"`
}

type SubUserCreateRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  *string `json:"lastName" binding:"omitempty"`
	Phone     string  `json:"phone" binding:"required"`
	Email     *string `json:"email" binding:"omitempty"`
	Role      int     `json:"role" binding:"required"`
}
