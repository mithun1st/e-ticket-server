package enum

type SuperUserRoleType int

const (
	Owner SuperUserRoleType = iota + 1
	Manager
)

type SubUserRoleType int

const (
	VehicleOwner SubUserRoleType = iota + 1
	CounterManager
	// Passanger
)

type VehicleType int

const (
	NonAcBus VehicleType = iota + 1
	AcBus
)

type PermissionType int

const (
	Hide PermissionType = iota + 1
	View
	Required
)
