package enum

type UserRoleType int

const (
	CompanyOwner UserRoleType = iota + 1
	BusOwner
	CounterManager
	Passanger
)

type VehicleType int

const (
	AcBus    VehicleType = 1
	NonAcBus VehicleType = 2
)

type Permission int

const (
	Hide Permission = iota + 1
	View
	Required
)
