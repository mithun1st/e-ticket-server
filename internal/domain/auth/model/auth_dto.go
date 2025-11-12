package authmodel

type AuthCompanyRequest struct {
	CompanyId int    `json:"companyId" binding:"required"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone" binding:"omitempty"`
	Password  string `json:"password" binding:"required"`
}

type AuthSubUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}
