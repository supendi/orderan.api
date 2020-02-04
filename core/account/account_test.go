package account_test

import (
	"context"
	"testing"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/account/inmem"
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
	if err == nil {
		t.Fatal("Should return err caused by duplicate email")
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
	if err == nil {
		t.Fatal("Should return err caused by duplicate phone")
	}
}
