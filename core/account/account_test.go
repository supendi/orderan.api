package account_test

import (
	"context"
	"testing"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/account/inmem"
	"github.com/supendi/orderan.api/pkg/errors"
)

var accountService *account.Service
var accountRepo *inmem.AccountRepository
var newContext = context.Background()

//GetAccountService singleton instance
func GetAccountService() *account.Service {
	if accountService == nil {
		hasher := account.NewBCryptHasher()
		accountService = account.NewAccountService(GetAccountRepo(), hasher)
	}
	return accountService
}

//GetAccountRepo singleton instance
func GetAccountRepo() *inmem.AccountRepository {
	if accountRepo == nil {
		accountRepo = inmem.NewAccountRepository([]*account.Account{})
	}
	return accountRepo
}

func TestRegisterAccount(t *testing.T) {
	registrant := &account.Registrant{
		Name:     "joe",
		Email:    "joe@gmail.com",
		Phone:    "0813",
		Password: "",
	}

	newAccount, err := GetAccountService().RegisterAccount(newContext, registrant)
	if err != nil {
		t.Fatal(err)
	}
	accountRecord, err := accountRepo.GetByEmail(newContext, newAccount.Email)
	if accountRecord == nil {
		t.Fatal("New account should be not nil")
	}
	if accountRecord.ID == "" {
		t.Fatal("New account ID should be not empty")
	}
}

func TestRegisterAccountFailDuplicateEmail(t *testing.T) {
	registrant1 := &account.Registrant{
		Name:     "Joe",
		Email:    "joe@gmail.com",
		Phone:    "0813",
		Password: "",
	}
	registrant2 := &account.Registrant{
		Name:     "John",
		Email:    "joe@gmail.com",
		Phone:    "0813",
		Password: "",
	}

	_, err := GetAccountService().RegisterAccount(newContext, registrant1)
	_, err = GetAccountService().RegisterAccount(newContext, registrant2)

	if !errors.IsAppError(err) {
		t.Fatal("Should return an AppError type")
	}
}

func TestRegisterAccountFailDuplicatePhone(t *testing.T) {
	registrant1 := &account.Registrant{
		Name:     "Joe",
		Email:    "joe@gmail.com",
		Phone:    "0813",
		Password: "",
	}
	registrant2 := &account.Registrant{
		Name:     "John",
		Email:    "john@gmail.com",
		Phone:    "0813",
		Password: "",
	}

	_, err := GetAccountService().RegisterAccount(newContext, registrant1)
	_, err = GetAccountService().RegisterAccount(newContext, registrant2)

	if !errors.IsAppError(err) {
		t.Fatal("Should return an AppError type")
	}
}

func TestUpdateAccount(t *testing.T) {
	registrant1 := &account.Registrant{
		Name:     "Andrew",
		Email:    "andrew@gmail.com",
		Phone:    "0815",
		Password: "myStrongPassword",
	}
	newAccount, err := GetAccountService().RegisterAccount(newContext, registrant1)
	if err != nil {
		t.Fatal(err)
	}

	updateRequest := account.UpdateRequest{
		ID:   newAccount.ID,
		Name: "Joe Haskell",
	}

	_, err = GetAccountService().UpdateAccount(newContext, updateRequest)
	if err != nil {
		t.Fatal(err)
	}

	existingAccount, err := GetAccountRepo().GetByID(newContext, newAccount.ID)
	if err != nil {
		t.Fatal(err)
	}
	if existingAccount.Name != updateRequest.Name {
		t.Fatalf("New account's name should be %s but got %s", updateRequest.Name, existingAccount.Name)
	}
	if existingAccount.Email != newAccount.Email {
		t.Fatalf("Account's email should not be changed. It must be %s but got %s", newAccount.Email, existingAccount.Email)
	}
	if existingAccount.Phone != newAccount.Phone {
		t.Fatalf("Account's Phone should not be changed. It must be %s but got %s", newAccount.Phone, existingAccount.Phone)
	}
	if existingAccount.Password != newAccount.Password {
		t.Fatalf("Account's password should not be changed. It must be %s but got %s", newAccount.Password, existingAccount.Password)
	}
}
