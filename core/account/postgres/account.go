package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/supendi/dbx"
	"github.com/supendi/orderan.api/core/account"
)

//AccountRepository implements account.Repository
type AccountRepository struct {
	db *dbx.Context
}

//NewAccountRepository returns a new Account Repository instance
func NewAccountRepository(dbContext *dbx.Context) *AccountRepository {
	return &AccountRepository{
		db: dbContext,
	}
}

//Add a new account into storage
func (me *AccountRepository) Add(ctx context.Context, newAccount *account.Account) (*account.Account, error) {
	newAccount.ID = uuid.New().String()

	statement := dbx.NewStatement("INSERT INTO account(id, name, email, phone, password, created_at) VALUES (:id, :name, :email, :phone, :password, NOW())")
	statement.AddParameter("id", newAccount.ID)
	statement.AddParameter("name", newAccount.Name)
	statement.AddParameter("email", newAccount.Email)
	statement.AddParameter("phone", newAccount.Phone)
	statement.AddParameter("password", newAccount.Password)

	me.db.AddStatement(statement)

	_, err := me.db.SaveChanges(ctx)

	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

//Update updates an existing account in storage
func (me *AccountRepository) Update(ctx context.Context, updatedAccount *account.Account) (*account.Account, error) {
	statement := dbx.NewStatement("UPDATE account SET name=:name, updated_at = NOW() WHERE id=:id")
	statement.AddParameter("id", updatedAccount.ID)
	statement.AddParameter("name", updatedAccount.Name)
	me.db.AddStatement(statement)
	_, err := me.db.SaveChanges(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

//GetByID returns an account by ID
func (me *AccountRepository) GetByID(ctx context.Context, id string) (*account.Account, error) {
	statement := dbx.NewStatement("SELECT * FROM account WHERE id = :id")
	statement.AddParameter("id", id)

	rows, err := me.db.QueryStatementContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var retrievedAccount *account.Account

	for rows.Next() {
		retrievedAccount = &account.Account{}
		err := rows.Scan(&retrievedAccount.ID, &retrievedAccount.Name, &retrievedAccount.Email, &retrievedAccount.Phone, &retrievedAccount.Password, &retrievedAccount.CreatedAt, &retrievedAccount.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return retrievedAccount, nil
	}

	return nil, nil
}

//GetByEmail returns an account by email
func (me *AccountRepository) GetByEmail(ctx context.Context, email string) (*account.Account, error) {
	statement := dbx.NewStatement("SELECT * FROM account WHERE email = :email")
	statement.AddParameter("email", email)

	rows, err := me.db.QueryStatementContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var retrievedAccount *account.Account

	for rows.Next() {
		retrievedAccount = &account.Account{}
		err := rows.Scan(&retrievedAccount.ID, &retrievedAccount.Name, &retrievedAccount.Email, &retrievedAccount.Phone, &retrievedAccount.Password, &retrievedAccount.CreatedAt, &retrievedAccount.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return retrievedAccount, nil
	}

	return nil, nil
}

//GetByPhone returns an account by phone
func (me *AccountRepository) GetByPhone(ctx context.Context, phone string) (*account.Account, error) {
	statement := dbx.NewStatement("SELECT * FROM account WHERE phone = :phone")
	statement.AddParameter("phone", phone)

	rows, err := me.db.QueryStatementContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var retrievedAccount *account.Account

	for rows.Next() {
		retrievedAccount = &account.Account{}
		err := rows.Scan(&retrievedAccount.ID, &retrievedAccount.Name, &retrievedAccount.Email, &retrievedAccount.Phone, &retrievedAccount.Password, &retrievedAccount.CreatedAt, &retrievedAccount.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return retrievedAccount, nil
	}

	return nil, nil
}
