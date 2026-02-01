package schedulemodel

import "time"

type ScheduleUri struct {
	CompanyId int `uri:"companyId" binding:"required"`
}

type ScheduleQuery struct {
	Date time.Time `form:"date" binding:"required"`
}
