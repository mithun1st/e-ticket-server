package homerepository

import (
	homemodel "e-ticket/internal/domain/home/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	"fmt"
)

type Repository struct {
	db *appdatabase.DbEntity
}

func NewHomeRepository(db *appdatabase.DbEntity) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindCompaniesByOwner(companyOwnerId int) ([]homemodel.CompaniesEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
	%s,
	%s,
	%s,
	%s,
	%s,
	%s
FROM %s WHERE(
	%s=%d
)
`,
		schema.Companies_id,
		schema.Companies_name,
		schema.Companies_address,
		schema.Companies_email,
		schema.Companies_phone,
		schema.Companies_is_active,

		schema.Companies,

		schema.CompaniesFkCompanyOwnerId, companyOwnerId,
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var companies []homemodel.CompaniesEntity

	for rows.Next() {
		var company homemodel.CompaniesEntity
		err := rows.Scan(
			&company.Id,
			&company.Name,
			&company.Address,
			&company.Email,
			&company.Phone,
			&company.IsActive,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (r *Repository) FindVehiclesByCompanies(companyIds []int) ([]homemodel.VehiclesEntity, error) {

	var sql string = fmt.Sprintf(`
// SELECT
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	%s,
// 	%s,
// 	%s,
// 	%s,
// 	%s.%s
// FROM %s
// LEFT JOIN %s ON
// 	%s.%s=%s.%s
// LEFT JOIN %s ON
// 	%s.%s=%s.%s
// WHERE %s IN (%s)
// `,
	// 		schema.VehiclesTableName, schema.Vehicle_id,
	// 		schema.Users, schema.Users_first_name,
	// 		schema.Users, schema.Users_last_name,
	// 		schema.Companies, schema.Companies_name,
	// 		schema.VehiclesTableName, schema.Vehicles_name,
	// 		schema.Vehicles_temporary_name,
	// 		schema.Vehicles_license_number,
	// 		schema.Vehicles_type,
	// 		schema.Vehicles_capacity,
	// 		schema.VehiclesTableName, schema.Vehicles_is_active,

	// 		schema.VehiclesTableName,

	// 		schema.Users,
	// 		schema.VehiclesTableName, schema.Vehicles_fk_owner_id,
	// 		schema.Users, schema.Users_id,

	// 		schema.Companies,
	// 		schema.VehiclesTableName, schema.Vehicles_fk_company_id,
	// 		schema.Companies, schema.Companies_id,

	// 		schema.Vehicles_fk_company_id,
	// 		utils.JoinArray(companyIds),
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var vehicles []homemodel.VehiclesEntity

	for rows.Next() {
		var vehicle homemodel.VehiclesEntity
		var firstName, lastname *string

		err := rows.Scan(
			&vehicle.Id,
			&firstName,
			&lastname,
			&vehicle.CompanyName,
			&vehicle.Name,
			&vehicle.TemporaryName,
			&vehicle.LicenseNumber,
			&vehicle.VehicleType,
			&vehicle.Capacity,
			&vehicle.IsActive,
		)
		if err != nil {
			return nil, err
		}
		if firstName != nil {
			vehicle.OwnerName = firstName

		}
		if lastname != nil {
			*vehicle.OwnerName = *vehicle.OwnerName + " " + *lastname
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}

func (r *Repository) FindOwnersByVehicles(vehiclesIds []int) ([]homemodel.OwnerEntity, error) {

	var sql string = fmt.Sprintf(`
// SELECT
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	%s.%s,
// 	count(%s.%s)
// FROM %s
// LEFT JOIN %s ON
// 	%s.%s=%s.%s
// WHERE %s.%s IN (%s)
// GROUP BY %s.%s`,
	// 		schema.Users, schema.Users_id,
	// 		schema.Users, schema.Users_first_name,
	// 		schema.Users, schema.Users_last_name,
	// 		schema.Users, schema.Users_email,
	// 		schema.Users, schema.Users_phone,
	// 		schema.Users, schema.Users_is_active,
	// 		schema.Users, schema.Users_id,

	// 		schema.VehiclesTableName,

	// 		schema.Users,
	// 		schema.VehiclesTableName, schema.Vehicles_fk_owner_id,
	// 		schema.Users, schema.Users_id,

	// 		schema.VehiclesTableName, schema.Vehicle_id,
	// 		utils.JoinArray(vehiclesIds),
	// 		schema.Users, schema.Users_id,
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var owners []homemodel.OwnerEntity

	for rows.Next() {
		var owner homemodel.OwnerEntity

		err := rows.Scan(
			&owner.Id,
			&owner.FirstName,
			&owner.LastName,
			&owner.Email,
			&owner.Phone,
			&owner.IsActive,
			&owner.NumOfVehicle,
		)
		if err != nil {
			return nil, err
		}

		owners = append(owners, owner)
	}

	return owners, nil
}

func (r *Repository) FindVehiclesByOwner(vehicleOwnerId int) ([]homemodel.VehiclesEntity, error) {

	var sql string = fmt.Sprintf(`
// SELECT
// %s.%s,
// %s.%s,
// %s.%s,
// %s,
// %s,
// %s,
// %s,
// %s.%s
// FROM %s
// LEFT JOIN %s
// ON %s.%s = %s.%s
// WHERE %s.%s = %d;
// `,
	// 		schema.VehiclesTableName, schema.Vehicle_id,
	// 		schema.Companies, schema.Companies_name,
	// 		schema.VehiclesTableName, schema.Vehicles_name,
	// 		schema.Vehicles_temporary_name,
	// 		schema.Vehicles_license_number,
	// 		schema.Vehicles_type,
	// 		schema.Vehicles_capacity,
	// 		schema.VehiclesTableName, schema.Vehicles_is_active,

	// 		schema.VehiclesTableName,
	// 		schema.Companies,

	// 		schema.VehiclesTableName, schema.Vehicles_fk_company_id,
	// 		schema.Companies, schema.Companies_id,

	// 		schema.VehiclesTableName, schema.Vehicles_fk_owner_id,
	// 		vehicleOwnerId,
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var vehicles []homemodel.VehiclesEntity

	for rows.Next() {
		var vehicle homemodel.VehiclesEntity

		err := rows.Scan(
			&vehicle.Id,
			&vehicle.CompanyName,
			&vehicle.Name,
			&vehicle.TemporaryName,
			&vehicle.LicenseNumber,
			&vehicle.VehicleType,
			&vehicle.Capacity,
			&vehicle.IsActive,
		)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}
