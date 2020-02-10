package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Needed for database test
	"github.com/supendi/orderan.api/core/env"
)

//NewDBTest return a new instance of sqlx.DB
func NewDBTest() *sqlx.DB {
	dsn := env.GetDBTestConstring()
	if dsn == "" {
		panic("DB connection string not found")
	}
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
