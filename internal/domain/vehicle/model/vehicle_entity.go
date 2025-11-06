package vehiclemodel

type VehicleEntity struct {
	Id            int     `json:"id"`
	OwnerId       int     `json:"ownerId"`
	CompanyId     int     `json:"companyId"`
	Name          string  `json:"name"`
	TemporaryName *string `json:"temporary_name"`
	LicenseNumber *string `json:"license_number"`
	VehicleType   int     `json:"type"`
	Capacity      int     `json:"capacity"`
	IsActive      bool    `json:"isActive"`
}
