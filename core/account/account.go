package account

import (
	"context"
	"time"

	"github.com/supendi/orderan.api/pkg/errors"
)

//Models
type (
	//Account represent account entity model
	Account struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		Phone     string     `json:"phone"`
		Password  string     `json:"-"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}

	//Registrant represent a registrant data model who wants to register as a new account
	Registrant struct {
		Name     string `validate:"required"`
		Email    string `validate:"required"`
		Phone    string `validate:"required"`
		Password string `validate:"required"`
	}

	//UpdateRequest represent account update request model
	UpdateRequest struct {
		AccountID string
		Name      string
	}

	//GetRequest represent account get request model
	GetRequest struct {
		AccountID string
	}
)

//Interfaces
type (
	//Repository specify the functionalities of account data storage
	Repository interface {
		Add(ctx context.Context, account *Account) (*Account, error)
		Update(ctx context.Context, account *Account) (*Account, error)
		GetByID(ctx context.Context, accountID string) (*Account, error)
		GetByEmail(ctx context.Context, email string) (*Account, error)
		GetByPhone(ctx context.Context, phone string) (*Account, error)
	}

	//PasswordHasher specify password hasher functions contract
	PasswordHasher interface {
		Hash(password string) (string, error)
		Verify(plainPassword string, hashedPassword string) (bool, error)
	}
)

//Service provide the account bussines functionalities such as create a new account, update and delete.
type Service struct {
	accountRepo    Repository
	passwordHasher PasswordHasher
}

//NewAccountService return new intance of account service
func NewAccountService(accountRepo Repository, passwordHasher PasswordHasher) *Service {
	return &Service{
		accountRepo:    accountRepo,
		passwordHasher: passwordHasher,
	}
}

//RegisterAccount registers a new account
func (me *Service) RegisterAccount(ctx context.Context, registrant *Registrant) (*Account, error) {
	existingAccount, err := me.accountRepo.GetByEmail(ctx, registrant.Email)
	if err != nil {
		return nil, err
	}
	if existingAccount != nil {
		appErr := errors.NewAppError("Validation error(s) occured.")
		appErr.Errors.Add(errors.NewFieldError("email", "Email '"+registrant.Email+"' is already registered."))
		return nil, appErr
	}

	existingAccount, err = me.accountRepo.GetByPhone(ctx, registrant.Phone)
	if err != nil {
		return nil, err
	}
	if existingAccount != nil {
		appErr := errors.NewAppError("Validation error(s) occured.")
		appErr.Errors.Add(errors.NewFieldError("phone", "Phone '"+registrant.Phone+"' is already registered."))
		return nil, appErr
	}

	newAccount := &Account{
		Name:      registrant.Name,
		Email:     registrant.Email,
		Phone:     registrant.Phone,
		Password:  registrant.Password,
		CreatedAt: time.Now(),
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

//UpdateAccount updates an existing account, but only its name will be updated
func (me *Service) UpdateAccount(ctx context.Context, updateRequest *UpdateRequest) (*Account, error) {
	existingAccount, err := me.accountRepo.GetByID(ctx, updateRequest.AccountID)
	if err != nil {
		return nil, err
	}
	if existingAccount == nil {
		appErr := errors.NewAppError("Validation error(s) occured.")
		appErr.Errors.Add(errors.NewFieldError("accountId", "Account with id '"+updateRequest.AccountID+"' is not found"))
		return nil, appErr
	}

	existingAccount.Name = updateRequest.Name

	updatedAccount, err := me.accountRepo.Update(ctx, existingAccount)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}

//GetAccount gets an account by its ID
func (me *Service) GetAccount(ctx context.Context, request *GetRequest) (*Account, error) {
	existingAccount, err := me.accountRepo.GetByID(ctx, request.AccountID)
	if err != nil {
		return nil, err
	}
	if existingAccount == nil {
		appErr := errors.NewAppError("Validation error(s) occured.")
		appErr.Errors.Add(errors.NewFieldError("accountId", "Account with id '"+request.AccountID+"' is not found"))
		return nil, appErr
	}

	return existingAccount, nil
}
