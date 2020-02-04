package account_test

import (
	"context"
	"testing"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/account/inmem"
)

func TestRegisterAccount(t *testing.T) {
	registrant := &account.Registrant{
		Name:     "joe",
		Email:    "joe@gmail.com",
		Phone:    "0813",
		Password: "",
	}
	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	hasher := account.NewBCryptHasher()
	accountSvc := account.NewAccountService(accountRepo, hasher)
	newContext := context.Background()
	newAccount, err := accountSvc.RegisterAccount(newContext, registrant)
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
