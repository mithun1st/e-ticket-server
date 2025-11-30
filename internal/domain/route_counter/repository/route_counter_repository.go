package routecounterrepository

import (
	routecountermodel "e-ticket/internal/domain/route_counter/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	"e-ticket/pkg/utils"
	"fmt"
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
%s,
%s
FROM %s
WHERE
%s=%d
`,
		schema.Counter_id,
		schema.Counter_fk_company_id,
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
			&counter.FkCompanyId,
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
%s,
%s
FROM %s
WHERE
%s=%d
`,
		schema.Route_id,
		schema.Route_fk_company_id,
		schema.Route_name,
		schema.Route_note,
		schema.Route_is_active,

		schema.Route,

		schema.Route_fk_company_id, companyId)

	rows, err := r.db.PQ.Query(sql)
	if err != nil {
		return nil, err
	}

	var routes []routecountermodel.RouteEntity

	for rows.Next() {
		var route routecountermodel.RouteEntity
		err := rows.Scan(
			&route.Id,
			&route.Fk_company_id,
			&route.Name,
			&route.Note,
			&route.Is_active,
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
