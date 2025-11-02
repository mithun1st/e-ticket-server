package appdatabase

import "database/sql"

type DbEntity struct {
	PQ *sql.DB
	MS *sql.DB
}
