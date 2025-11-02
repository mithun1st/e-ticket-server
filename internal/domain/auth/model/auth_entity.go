package authmodel

type AuthEntity struct {
	Token      string     `json:"token" binding:"required,"`
	UserEntity UserEntity `json:"user"`
	CompanyId  *int       `json:"companyId"`
	Role       int        `json:"role"`
}
