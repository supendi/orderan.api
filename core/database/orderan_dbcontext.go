package database

import "github.com/supendi/dbx"

//OrderanDBContext represent orderan database
type OrderanDBContext struct {
	*dbx.Context
}

//NewOrderanDBContext return a new instance of orderan database context
func NewOrderanDBContext(dbContext *dbx.Context) *OrderanDBContext {
	return &OrderanDBContext{
		Context: dbContext,
	}
}

//TruncateTables truncate all tables in the database
func (me *OrderanDBContext) TruncateTables() error {
	statement := dbx.NewStatement(`
	  TRUNCATE TABLE
	  account,
	  token
	  `)
	me.AddStatement(statement)
	_, err := me.SaveChanges(nil)
	return err
}
