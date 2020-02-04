package account

import (
	"context"
	"time"
)

//Account represent account entity model
type Account struct {
	ID        string
	Name      string
	Email     string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

//Registrant represent a registrant data model who wants to register as a new account
type Registrant struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

//Repository specify the functionalities of account data storage
type Repository interface {
	Add(account *Account) (*Account, error)
}

//PasswordHasher specify password hasher functions contract
type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(plainPassword string, hashedPassword string) (bool, error)
}

//Service provide the account bussines functionalities such as create a new account, update and delete.
type Service struct {
	accountRepo    Repository
	passwordHasher PasswordHasher
}

//RegisterAccount registers a new account
func (me *Service) RegisterAccount(ctx context.Context, registrant *Registrant) (*Account, error) {
	newAccount := &Account{
		Name:     registrant.Name,
		Email:    registrant.Email,
		Phone:    registrant.Phone,
		Password: registrant.Password,
	}

	hashedPassword, err := me.passwordHasher.Hash(newAccount.Password)
	if err != nil {
		return nil, err
	}
	newAccount.Password = hashedPassword

	addedAccount, err := me.accountRepo.Add(newAccount)
	if err != nil {
		return nil, err
	}

	return addedAccount, nil
}
