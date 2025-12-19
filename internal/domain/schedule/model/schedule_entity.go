package schedulemodel

import "time"

type ScheduleEntity struct {
	Id           int       `json:"id"`
	FkRouteId    int       `json:"fkRouteId"`
	RouteName    string    `json:"routeName"`
	FkVehicleId  int       `json:"fkVehicleId"`
	VehicleName  string    `json:"vehicleName"`
	StartAt      time.Time `json:"startAt"`
	RepeatsDaily bool      `json:"repeatsDaily"`
}
