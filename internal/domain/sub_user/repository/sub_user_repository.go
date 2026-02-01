package subuserrepository

import (
	subusermodel "e-ticket/internal/domain/sub_user/model"
	"e-ticket/internal/schema"
	appdatabase "e-ticket/pkg/database"
	appenviroment "e-ticket/pkg/enviroment"
	applogger "e-ticket/pkg/logger"
	"e-ticket/pkg/utils"
	"fmt"
)

type Repository struct {
	db *appdatabase.DbEntity
}

func NewSubUserRepository(db *appdatabase.DbEntity) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUsersIdBy(companyId int, role *int) ([]int, error) {
	var sql string

	sql = fmt.Sprintf(`
	SELECT %s 
	FROM %s WHERE
	%s=%d
	`,
		schema.CompanySubUsers_user_id,
		schema.CompanySubUsers,
		schema.CompanySubUsers_company_id,
		companyId,
	)

	if role != nil {
		str := fmt.Sprintf(`
AND
%s=%d
`,
			schema.CompanySubUsers_role,
			*role,
		)
		sql = sql + str
	}

	if appenviroment.GetCuerrentStatus() == appenviroment.Development {
		applogger.Info(sql)
	}
	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}
	var userIds []int
	for rows.Next() {
		var id int
		err := rows.Scan(
			&id,
		)
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, id)
	}
	return userIds, nil
}

func (r *Repository) FindUsersByIds(userIds []int) ([]subusermodel.UserEntity, error) {

	var sql string = fmt.Sprintf(`
SELECT
%s,
%s,
%s,
%s,
%s,
%s
FROM %s WHERE
%s IN (%s)
`,
		schema.Users_id,
		schema.Users_first_name,
		schema.Users_last_name,
		schema.Users_phone,
		schema.Users_email,
		schema.Users_is_active,

		schema.Users,

		schema.Users_id,
		utils.DbValues(
			utils.ModelsToElements(
				userIds,
				func(id int) any {
					return id
				},
			)...,
		),
	)

	if appenviroment.GetCuerrentStatus() == appenviroment.Development {
		applogger.Info(sql)
	}
	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}
	var users []subusermodel.UserEntity
	for rows.Next() {
		var user subusermodel.UserEntity
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
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) FindUserByPhone(phone string) (*subusermodel.UserEntity, error) {
	var sql string = fmt.Sprintf(`
SELECT
%s,
%s,
%s,
%s,
%s,
%s
FROM %s WHERE
%s = '%s'
`,
		schema.Users_id,
		schema.Users_first_name,
		schema.Users_last_name,
		schema.Users_phone,
		schema.Users_email,
		schema.Users_is_active,

		schema.Users,

		schema.Users_phone,
		phone,
	)

	if appenviroment.GetCuerrentStatus() == appenviroment.Development {
		applogger.Info(sql)
	}
	rows, err := r.db.PQ.Query(sql)

	if err != nil {
		return nil, err
	}

	var user subusermodel.UserEntity
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

func (r *Repository) InsertUser(
	firstName string,
	lastName *string,
	phone string,
	email *string,
	password string,
) (bool, error) {
	var sql string = fmt.Sprintf(`
INSERT INTO %s (
	%s,
	%s,
	%s,
	%s,
	%s
) VALUES (
	'%s',
	%s,
	'%s',
	%s,
	'%s'
)`,
		schema.Users,

		schema.Users_first_name,
		schema.Users_last_name,
		schema.Users_phone,
		schema.Users_email,
		schema.Users_password_hash,

		firstName,
		utils.DbNilToStr(lastName),
		phone,
		utils.DbNilToStr(email),
		password,
	)

	if appenviroment.GetCuerrentStatus() == appenviroment.Development {
		applogger.Info(sql)
	}
	result, err := r.db.PQ.Exec(sql)

	if err != nil {
		return false, err
	}

	rowsAffects, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffects != 0, nil
}

func (r *Repository) InsertCompanySubUser(
	companyId int,
	userId int,
	role int,
) (bool, error) {
	var sql string = fmt.Sprintf(`
INSERT INTO %s (
	%s,
	%s,
	%s
) VALUES (
	%d,
	%d,
	%d
)`,
		schema.CompanySubUsers,

		schema.CompanySubUsers_company_id,
		schema.CompanySubUsers_user_id,
		schema.CompanySubUsers_role,

		companyId,
		userId,
		role,
	)

	if appenviroment.GetCuerrentStatus() == appenviroment.Development {
		applogger.Info(sql)
	}
	result, err := r.db.PQ.Exec(sql)

	if err != nil {
		return false, err
	}

	rowsAffects, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffects != 0, nil
}
