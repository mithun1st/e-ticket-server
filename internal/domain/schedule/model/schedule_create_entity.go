package schedulemodel

import "time"

type ScheduleCreateEntity struct {
	FkRouteId   int       `json:"fkRouteId" binding:"required"`
	FkVehicleId int       `json:"fkVehicleId" binding:"required"`
	StartAt     time.Time `json:"startAt" binding:"required"`
}
