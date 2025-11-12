package vehiclerepository

import (
	vehiclemodel "e-ticket/internal/domain/vehicle/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	"e-ticket/pkg/utils"
	"fmt"
)

type Repository struct {
	db *appdatabase.DbEntity
}

func NewVehicleRepository(db *appdatabase.DbEntity) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindVehiclesByCompanyAndOwner(companyId int, userId *int) ([]vehiclemodel.VehicleEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
%s,
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
		schema.Vehicles_id,
		schema.Vehicles_fk_owner_id,
		schema.Vehicles_fk_company_id,
		schema.Vehicles_name,
		schema.Vehicles_temporary_name,
		schema.Vehicles_license_number,
		schema.Vehicles_type,
		schema.Vehicles_capacity,
		schema.Vehicles_is_active,
		schema.Vehicles,
		schema.Vehicles_fk_company_id, companyId,
	)

	if userId != nil {
		str := fmt.Sprintf(`
AND
%s=%d`,
			schema.Vehicles_fk_owner_id, *userId)
		sql = sql + str
	}

	rows, err := r.db.PQ.Query(sql)
	if err != nil {
		return nil, err
	}
	var list []vehiclemodel.VehicleEntity
	for rows.Next() {
		var vehicleEntity vehiclemodel.VehicleEntity
		err := rows.Scan(
			&vehicleEntity.Id,
			&vehicleEntity.OwnerId,
			&vehicleEntity.CompanyId,
			&vehicleEntity.Name,
			&vehicleEntity.TemporaryName,
			&vehicleEntity.LicenseNumber,
			&vehicleEntity.VehicleType,
			&vehicleEntity.Capacity,
			&vehicleEntity.IsActive,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, vehicleEntity)
	}
	return list, nil
}

func (r *Repository) InsertVehicle(companyId int, vehicle vehiclemodel.VehicleCreateRequest) (bool, error) {

	var sql string = fmt.Sprintf(`
INSERT INTO %s(
%s,
%s,
%s,
%s,
%s,
%s,
%s
) VALUES (
%d,
%d,
'%s',
%s,
%s,
%d,
%d	
)
`,
		schema.Vehicles,

		schema.Vehicles_fk_owner_id,
		schema.Vehicles_fk_company_id,

		schema.Vehicles_name,
		schema.Vehicles_temporary_name,
		schema.Vehicles_license_number,
		schema.Vehicles_type,
		schema.Vehicles_capacity,

		vehicle.OwnerId,
		companyId,

		vehicle.Name,
		utils.NilToStrDB(vehicle.TemporaryName),
		utils.NilToStrDB(vehicle.LicenseNumber),
		vehicle.VehicleType,
		vehicle.Capacity,
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
