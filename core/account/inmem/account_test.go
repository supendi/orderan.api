package inmem

import (
	"context"
	"testing"

	"github.com/supendi/orderan.api/core/account"
)

func TestAddAccount(t *testing.T) {
	registrant := &account.Account{
		Name:     "Irpan",
		Email:    "irpan.nowan@gmail.com",
		Phone:    "+1234",
		Password: "strongPassw04rd",
	}
	newContext := context.Background()
	repo := NewAccountRepository([]*account.Account{})
	_, err := repo.Add(newContext, registrant)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.GetByEmail(newContext, registrant.Email)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("Retrieved account must be not nil")
	}
	if got.Email != registrant.Email {
		t.Fatal("Wrong email")
	}
	if got.ID == "" {
		t.Fatal("I want a non empty ID")
	}
	t.Log(got.ID)
}

func TestGetAccountByEmail(t *testing.T) {
	registrant := &account.Account{
		Name:     "Irpan",
		Email:    "irpan.nowan@gmail.com",
		Phone:    "+1234",
		Password: "strongPassw04rd",
	}
	newContext := context.Background()
	repo := NewAccountRepository([]*account.Account{})
	_, err := repo.Add(newContext, registrant)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.GetByEmail(newContext, registrant.Email)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("Retrieved account must be not nil")
	}
	if got.Email != registrant.Email {
		t.Fatal("Wrong email")
	}
	if got.ID == "" {
		t.Fatal("I want a non empty ID")
	}
	t.Log(got.ID)
}

func TestGetAccountByPhone(t *testing.T) {
	registrant := &account.Account{
		Name:     "Irpan",
		Email:    "irpan.nowan@gmail.com",
		Phone:    "+1234",
		Password: "strongPassw04rd",
	}
	newContext := context.Background()
	repo := NewAccountRepository([]*account.Account{})
	_, err := repo.Add(newContext, registrant)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.GetByPhone(newContext, registrant.Phone)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("Retrieved account must be not nil")
	}
	if got.Phone != registrant.Phone {
		t.Fatal("Wrong phone")
	}
	if got.ID == "" {
		t.Fatal("I want a non empty ID")
	}
	t.Log(got.ID)
}
func TestGetAccountByID(t *testing.T) {
	registrant := &account.Account{
		Name:     "Lady",
		Email:    "lady.gaga@gmail.com",
		Phone:    "+1234",
		Password: "strongPassw04rd",
	}
	newContext := context.Background()
	repo := NewAccountRepository([]*account.Account{})
	newAccount, err := repo.Add(newContext, registrant)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.GetByID(newContext, newAccount.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("Retrieved account must be not nil")
	}
	if got.ID != newAccount.ID {
		t.Fatal("Wrong ID")
	}
	t.Log(got.ID)
}

func TestUpdateAccount(t *testing.T) {
	registrant := &account.Account{
		Name:     "Lady",
		Email:    "lady.gaga@gmail.com",
		Phone:    "+1234",
		Password: "strongPassw04rd",
	}
	newContext := context.Background()
	repo := NewAccountRepository([]*account.Account{})
	newAccount, err := repo.Add(newContext, registrant)
	if err != nil {
		t.Fatal(err)
	}
	existingAccount, err := repo.GetByID(newContext, newAccount.ID)
	if err != nil {
		t.Fatal(err)
	}

	if existingAccount == nil {
		t.Fatal("Retrieved account must be not nil")
	}
	existingAccount.Name = "newName"
	updatedAccount, err := repo.Update(newContext, existingAccount)
	if err != nil {
		t.Fatal(err)
	}
	if updatedAccount.Name != "newName" {
		t.Fatal("Wrong Name")
	}
}
