package inmem

import (
	"context"

	"github.com/google/uuid"
	"github.com/supendi/orderan.api/core/account"
)

//AccountRepository implement account.AccountRepository which use memory as data storage
type AccountRepository struct {
	accounts []*account.Account
}

//NewAccountRepository returns a new Account Repository instance
func NewAccountRepository(accounts []*account.Account) *AccountRepository {
	return &AccountRepository{
		accounts: accounts,
	}
}

//Add a new account to memory
func (me *AccountRepository) Add(ctx context.Context, newAccount *account.Account) (*account.Account, error) {
	newAccount.ID = uuid.New().String()
	me.accounts = append(me.accounts, newAccount)
	return newAccount, nil
}

//Update updates an existing account. NOT including its email, password and phone
func (me *AccountRepository) Update(ctx context.Context, updatedAccount *account.Account) (*account.Account, error) {
	for _, accountRecord := range me.accounts {
		if accountRecord.ID == updatedAccount.ID {
			accountRecord.Name = updatedAccount.Name
			return accountRecord, nil
		}
	}
	return nil, nil
}

//GetByID gets an account by its ID
func (me *AccountRepository) GetByID(ctx context.Context, id string) (*account.Account, error) {
	for _, accountRecord := range me.accounts {
		if accountRecord.ID == id {
			return accountRecord, nil
		}
	}
	return nil, nil
}

//GetByEmail gets an account by its email
func (me *AccountRepository) GetByEmail(ctx context.Context, email string) (*account.Account, error) {
	for _, accountRecord := range me.accounts {
		if accountRecord.Email == email {
			return accountRecord, nil
		}
	}
	return nil, nil
}

//GetByPhone gets an account by its phone numer
func (me *AccountRepository) GetByPhone(ctx context.Context, phone string) (*account.Account, error) {
	for _, accountRecord := range me.accounts {
		if accountRecord.Phone == phone {
			return accountRecord, nil
		}
	}
	return nil, nil
}
