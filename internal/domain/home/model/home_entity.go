package homemodel

type CompaniesEntity struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Address  *string `json:"address"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	IsActive bool    `json:"isActive"`
}

type OwnerEntity struct {
	Id           int     `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     *string `json:"lastName"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	IsActive     bool    `json:"isActive"`
	NumOfVehicle int     `json:"numOfVehicle"`
}

type VehiclesEntity struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	TemporaryName *string `json:"temporary_name"`
	LicenseNumber *string `json:"license_number"`
	VehicleType   int     `json:"type"`
	Capacity      int     `json:"capacity"`
	IsActive      bool    `json:"isActive"`
	OwnerName     *string `json:"ownerName"`
	CompanyName   string  `json:"companyName"`
}

type HomeEntity struct {
	Companies *[]CompaniesEntity `json:"companies"`
	Owners    *[]OwnerEntity     `json:"owners"`
	Vehicles  *[]VehiclesEntity  `json:"vehicles"`
}
