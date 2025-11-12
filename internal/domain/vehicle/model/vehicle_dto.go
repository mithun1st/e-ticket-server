package vehiclemodel

type VehicleUri struct {
	CompanyId int `uri:"companyId" binding:"required"`
}

type VehicleQuery struct {
	UserId *int `form:"userId" binding:"omitempty"`
}
type VehicleCreateRequest struct {
	OwnerId       int     `json:"ownerId" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	TemporaryName *string `json:"temporary_name" binding:"omitempty"`
	LicenseNumber *string `json:"license_number" binding:"omitempty"`
	VehicleType   int     `json:"type" binding:"required"`
	Capacity      int     `json:"capacity" binding:"required"`
}
