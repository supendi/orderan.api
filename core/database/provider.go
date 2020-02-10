package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Needed for database test
	"github.com/supendi/dbx"
	"github.com/supendi/orderan.api/core/env"
)

//dsn := env.GetDBTestConstring()

//newDB return a new instance of sqlx.DB
func newDB(connectioString string) *sqlx.DB {
	if connectioString == "" {
		panic("DB connection string must not an empty string")
	}
	db, err := sqlx.Open("postgres", connectioString)
	if err != nil {
		panic(err)
	}
	return db
}

//NewDBContext return a new dbxContext
func newDBContext(db *sqlx.DB) *dbx.Context {
	dbClient := dbx.NewClient(db)
	dbContext := dbx.NewContext(dbClient)
	return dbContext
}

//NewDBContextTest return a new orderan db context test
func NewDBContextTest() *dbx.Context {
	constring := env.GetDBTestConstring()
	db := newDB(constring)
	dbContext := newDBContext(db)
	return dbContext
}

//NewDBContext return a new orderan db context
func NewDBContext() *dbx.Context {
	constring := env.GetDBConstring()
	db := newDB(constring)
	dbContext := newDBContext(db)
	return dbContext
}
