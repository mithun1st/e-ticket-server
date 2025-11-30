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

func (r *Repository) InsertCounter(companyId int, counters routecountermodel.CounterCreateEntity) (bool, error) {
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
		utils.DbNilToStr(counters.FkAssignUserId),
		counters.Name,
		utils.DbNilToStr(counters.Address),
		utils.DbNilToStr(counters.Lat),
		utils.DbNilToStr(counters.Long),
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

// func (r *Repository) FindAll() ([]routecountermodel.RouteCounterEntity, error) {

// 	var sql string = fmt.Sprintln(`select * from db`)

// 	rows, err := r.db.PQ.Query(sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var list []routecountermodel.RouteCounterEntity
// 	for rows.Next() {
// 		var routeCounterEntity routecountermodel.RouteCounterEntity
// 		err := rows.Scan(
// 			&routeCounterEntity.Id,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		list = append(list, routeCounterEntity)
// 	}
// 	return list, nil
// }
