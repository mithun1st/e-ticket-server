package appdatabase

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresDb(dbConnectionString string) (*sql.DB, error) {

	postgres, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}
	return postgres, nil

}

func Check(db *sql.DB) error {
	err := db.Ping()
	return err
}

func Close(db *sql.DB) error {
	err := db.Close()
	return err
}
