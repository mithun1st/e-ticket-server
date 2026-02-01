package enum

type superUserRoleType int

const (
	Owner superUserRoleType = iota + 1
	Manager
)

type subUserRoleType int

const (
	VehicleOwner subUserRoleType = iota + 1
	CounterManager
	Driver
	Supervisor
)

type vehicleType int

const (
	NonAcBus vehicleType = iota + 1
	AcBus
)

type permissionType int

const (
	Hide permissionType = iota + 1
	View
	Required
)
