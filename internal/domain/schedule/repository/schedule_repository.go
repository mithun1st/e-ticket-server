package schedulerepository

import (
	schedulemodel "e-ticket/internal/domain/schedule/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	"e-ticket/pkg/utils"
)

type Repository struct {
	db *appdatabase.DbEntity
}

func NewScheduleRepository(db *appdatabase.DbEntity) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindSchedulesById(companyId int) (*[]schedulemodel.ScheduleEntity, error) {

	var sql string = "WITH T1 AS ( SELECT " +
		utils.DbNames(
			schema.Route_id,
			schema.Route_name,
		) +
		" FROM " + schema.Route +
		" WHERE " + schema.Route_fk_company_id + "=" + utils.DbValues(companyId) +
		" ), T2 AS ( " +
		" SELECT " +
		utils.DbNames(
			schema.Vehicles_id,
			schema.Vehicles_name,
		) +
		" FROM " + schema.Vehicles +
		" WHERE " + schema.Vehicles_fk_company_id + "=" + utils.DbValues(companyId) +
		" ) " +
		"SELECT " +
		utils.DbNames(
			schema.Schedule+"."+schema.Schedule_id,
			schema.Schedule_fk_route_id,
			schema.Schedule_fk_vehicle_id,
			schema.Schedule_start_at,
			schema.Schedule_repeats_daily,
			"T1."+schema.Route_name,
			"T2."+schema.Vehicles_name,
		) +
		" FROM " + schema.Schedule +
		" LEFT JOIN T1 ON " + schema.Schedule_fk_route_id + " = T1." + schema.Route_id +
		" LEFT JOIN T2 ON " + schema.Schedule_fk_vehicle_id + " = T2." + schema.Vehicles_id +
		" WHERE " + schema.Schedule_fk_company_id +
		" = " + utils.DbValues(companyId) +
		" ORDER BY " + schema.Schedule_start_at +
		" ASC"

	rows, err := r.db.PQ.Query(sql)
	if err != nil {
		return nil, err
	}
	var schedules []schedulemodel.ScheduleEntity
	for rows.Next() {
		var schedule schedulemodel.ScheduleEntity
		err := rows.Scan(
			&schedule.Id,
			&schedule.FkRouteId,
			&schedule.FkVehicleId,
			&schedule.StartAt,
			&schedule.RepeatsDaily,
			&schedule.RouteName,
			&schedule.VehicleName,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return &schedules, nil
}

func (r *Repository) InsertSchedule(companyId int, schedule schedulemodel.ScheduleCreateEntity) (bool, error) {

	var sql string = "INSERT INTO " + schema.Schedule +
		" (" + utils.DbNames(
		schema.Schedule_fk_company_id,
		schema.Schedule_fk_route_id,
		schema.Schedule_fk_vehicle_id,
		schema.Schedule_start_at,
		schema.Schedule_repeats_daily,
	) + ") VALUES (" + utils.DbValues(
		companyId,
		schedule.FkRouteId,
		schedule.FkVehicleId,
		schedule.StartAt,
		schedule.RepeatsDaily,
	) + ")"

	result, err := r.db.PQ.Exec(sql)
	if err != nil {
		return false, err
	}

	rowEffect, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowEffect != 0, nil
}
