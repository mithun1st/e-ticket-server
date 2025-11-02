package authrepository

import (
	authmodel "e-ticket/internal/domain/auth/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	"fmt"
)

type Repository struct {
	db *appdatabase.DbEntity
}

func NewAuthRepository(db *appdatabase.DbEntity) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUser(email string, phone string, password string) (*authmodel.UserEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
	%s,
	%s,
	%s,
	%s,
	%s,
	%s
FROM %s WHERE(
	(%s='%s' OR
	%s='%s') AND
	%s='%s'
) LIMIT 1
`,
		schema.Users_id,
		schema.Users_first_name,
		schema.Users_last_name,
		schema.Users_phone,
		schema.Users_email,
		schema.Users_is_active,

		schema.Users,

		schema.Users_email, email,
		schema.Users_phone, phone,
		schema.Users_password_hash, password,
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var user authmodel.UserEntity

	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Phone,
			&user.Email,
			&user.IsActive,
		)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r *Repository) FindCompanyUser(ownerId int) (*int, *int, error) {

	var sql string = fmt.Sprintf(`
SELECT
	%s,
	%s
FROM %s WHERE(
	%s=%d
) LIMIT 1
`,
		schema.CompanyUsers_company_id,
		schema.CompanyUsers_role,
		schema.CompanyUsers,
		schema.CompanyUsers_user_id,
		ownerId,
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, nil, err
	}

	var companyId *int
	var role *int

	for rows.Next() {
		err := rows.Scan(
			&companyId,
			&role,
		)
		if err != nil {
			return nil, nil, err
		}
	}

	return companyId, role, nil
}

func (r *Repository) FindSubUserRole(userId int, ownerId int) (*int, error) {

	var sql string = fmt.Sprintf(`
SELECT
	%s
FROM %s WHERE(
	%s=%d AND
	%s=%d
) LIMIT 1
`,
		schema.CompanySubUsers_role,
		schema.CompanySubUsers,
		schema.CompanySubUsers_user_id,
		userId,
		schema.CompanySubUsers_company_id,
		ownerId,
	)

	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var role *int

	for rows.Next() {
		err := rows.Scan(
			&role,
		)
		if err != nil {
			return nil, err
		}
	}

	return role, nil
}
