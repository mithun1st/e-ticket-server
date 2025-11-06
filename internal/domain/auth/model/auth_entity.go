package authmodel

import subusermodel "e-ticket/internal/domain/sub_user/model"

type AuthEntity struct {
	Token      string                  `json:"token" binding:"required,"`
	UserEntity subusermodel.UserEntity `json:"user"`
	CompanyId  *int                    `json:"companyId"`
	Role       int                     `json:"role"`
}
