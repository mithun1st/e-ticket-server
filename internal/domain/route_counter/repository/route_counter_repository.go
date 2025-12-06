package routecounterrepository

import (
	routecountermodel "e-ticket/internal/domain/route_counter/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	"e-ticket/pkg/utils"
	"fmt"
	"strings"
)

type Repository struct {
	db *appdatabase.DbEntity
}

func NewRouteCounterRepository(db *appdatabase.DbEntity) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindCountersByCompanyId(companyId int) ([]routecountermodel.CounterEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
%s,
%s,
%s,
%s,
%s,
%s,
%s
FROM %s
WHERE
%s=%d
`,
		schema.Counter_id,
		schema.Counter_fk_assign_user_id,
		schema.Counter_name,
		schema.Counter_address,
		schema.Counter_lat,
		schema.Counter_long,
		schema.Counter_is_active,

		schema.Counter,

		schema.Counter_fk_company_id, companyId)

	rows, err := r.db.PQ.Query(sql)
	if err != nil {
		return nil, err
	}

	var counters []routecountermodel.CounterEntity

	for rows.Next() {
		var counter routecountermodel.CounterEntity
		err := rows.Scan(
			&counter.Id,
			&counter.FkAssignUserId,
			&counter.Name,
			&counter.Address,
			&counter.Lat,
			&counter.Long,
			&counter.IsActive,
		)
		if err != nil {
			return nil, err
		}

		counters = append(counters, counter)
	}

	return counters, nil
}

func (r *Repository) InsertCounter(companyId int, counter routecountermodel.CounterCreateEntity) (bool, error) {
	var sql string = fmt.Sprintf(`
INSERT INTO %s(
%s,
%s,
%s,
%s,
%s,
%s
) VALUES (
%d,
%s,
'%s',
%s,
%s,
%s
)`,
		schema.Counter,

		schema.Counter_fk_company_id,
		schema.Counter_fk_assign_user_id,
		schema.Counter_name,
		schema.Counter_address,
		schema.Counter_lat,
		schema.Counter_long,

		companyId,
		utils.DbNilToStr(counter.FkAssignUserId),
		counter.Name,
		utils.DbNilToStr(counter.Address),
		utils.DbNilToStr(counter.Lat),
		utils.DbNilToStr(counter.Long),
	)

	resutl, err := r.db.PQ.Exec(sql)
	if err != nil {
		return false, err
	}

	rowAffects, err := resutl.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowAffects != 0, nil
}

func (r *Repository) FindRoutesByCompanyId(companyId int) ([]routecountermodel.RouteEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
%s,
%s,
%s,
%s
FROM %s
WHERE
%s=%d
`,
		schema.Route_id,
		schema.Route_name,
		schema.Route_note,
		schema.Route_is_active,

		schema.Route,

		schema.Route_fk_company_id, companyId,
	)

	rows, err := r.db.PQ.Query(sql)
	if err != nil {
		return nil, err
	}

	var routes []routecountermodel.RouteEntity

	for rows.Next() {
		var route routecountermodel.RouteEntity
		err := rows.Scan(
			&route.Id,
			&route.Name,
			&route.Note,
			&route.IsActive,
		)
		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	return routes, nil
}

func (r *Repository) InsertRoute(companyId int, route routecountermodel.RouteCreateEntity) (bool, error) {
	var sql string = fmt.Sprintf(`
INSERT INTO %s(
%s,
%s,
%s
) VALUES (
%d,
'%s',
%s
)`,
		schema.Route,

		schema.Route_fk_company_id,
		schema.Route_name,
		schema.Route_note,

		companyId,
		route.Name,
		utils.DbNilToStr(route.Note),
	)

	resutl, err := r.db.PQ.Exec(sql)
	if err != nil {
		return false, err
	}

	rowAffects, err := resutl.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowAffects != 0, nil
}

func (r *Repository) FindRouteCountersByCompanyAndRoute(companyId int, routeId int) ([]routecountermodel.RouteCounterEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
%s,
%s,
%s,
%s
FROM %s
WHERE
%s=%d AND
%s=%d
ORDER BY %s ASC
`,
		schema.RouteCounter_counter_id,
		schema.RouteCounter_duration,
		schema.RouteCounter_cost,
		schema.RouteCounter_serial,

		schema.RouteCounter,

		schema.RouteCounter_fk_company_id, companyId,
		schema.RouteCounter_route_id, routeId,

		schema.RouteCounter_serial,
	)

	rows, err := r.db.PQ.Query(sql)
	if err != nil {
		return nil, err
	}

	var routeCounters []routecountermodel.RouteCounterEntity

	for rows.Next() {
		var routeCounter routecountermodel.RouteCounterEntity
		err := rows.Scan(
			&routeCounter.CounterId,
			&routeCounter.Duration,
			&routeCounter.Cost,
			&routeCounter.Serial,
		)
		if err != nil {
			return nil, err
		}

		routeCounters = append(routeCounters, routeCounter)
	}

	return routeCounters, nil
}

func (r *Repository) InsertRouteCounter(companyId int, routeId int, routeCounters []routecountermodel.RouteCounterCreateEntity) (bool, error) {

	var values []string
	for _, e := range routeCounters {
		var str string = fmt.Sprintf(`(%d, %d, %d, %s, %s, %d)`,
			routeId,
			e.CounterId,
			1,
			utils.DbNilToStr(e.Duration),
			utils.DbNilToStr(e.Cost),
			e.Serial,
		)
		values = append(values, str)
	}
	var valueAsStr string = strings.Join(values, ",")

	var sql string = fmt.Sprintf(`
INSERT INTO %s(
%s,
%s,
%s,
%s,
%s,
%s
) VALUES
%s
`,
		schema.RouteCounter,

		schema.RouteCounter_route_id,
		schema.RouteCounter_counter_id,
		schema.RouteCounter_fk_company_id,
		schema.RouteCounter_duration,
		schema.RouteCounter_cost,
		schema.RouteCounter_serial,

		valueAsStr,
	)

	fmt.Println(sql)

	resutl, err := r.db.PQ.Exec(sql)
	if err != nil {
		return false, err
	}

	rowAffects, err := resutl.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowAffects != 0, nil
}
