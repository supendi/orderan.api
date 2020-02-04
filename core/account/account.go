package account

import (
	"context"
	"time"

	"github.com/supendi/orderan.api/pkg"
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
	Add(ctx context.Context, account *Account) (*Account, error)
	GetByEmail(ctx context.Context, email string) (*Account, error)
	GetByPhone(ctx context.Context, phone string) (*Account, error)
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
	existingAccount, err := me.accountRepo.GetByEmail(ctx, registrant.Email)
	if err != nil {
		return nil, err
	}
	if existingAccount != nil {
		return nil, pkg.NewAppError("Email '" + registrant.Email + "' is already registered.")
	}

	existingAccount, err = me.accountRepo.GetByPhone(ctx, registrant.Phone)
	if err != nil {
		return nil, err
	}
	if existingAccount != nil {
		return nil, pkg.NewAppError("Phone number '" + registrant.Email + "' is already registered.")
	}

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

	addedAccount, err := me.accountRepo.Add(ctx, newAccount)
	if err != nil {
		return nil, err
	}

	return addedAccount, nil
}
