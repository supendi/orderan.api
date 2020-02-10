package postgres

import (
	"context"
	"testing"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/database"
)

func TestAccountAdd(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	accountRepo := NewAccountRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedAccount, err := accountRepo.Add(ctx, &account.Account{
		Name:     "Supendi Saja",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if insertedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	fetchedAccount, err := accountRepo.GetByID(ctx, insertedAccount.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if fetchedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	// err = orderanDBContext.TruncateTables()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	orderanDBContext.Close()
}

func TestAccountUpdate(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	accountRepo := NewAccountRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedAccount, err := accountRepo.Add(ctx, &account.Account{
		Name:     "Supendi Saja",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if insertedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	_, err = accountRepo.Update(ctx, &account.Account{
		ID:       insertedAccount.ID,
		Name:     "Supendi",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	fetchedAccount, err := accountRepo.GetByID(ctx, insertedAccount.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if fetchedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	if fetchedAccount.Name != "Supendi" {
		t.Fatalf("Updated account name must 'Supendi' but got %s", fetchedAccount.Name)
	}
	if fetchedAccount.Email != insertedAccount.Email {
		t.Fatalf("Updated account name must not modify the email. The email should be %s but got %s", insertedAccount.Email, fetchedAccount.Email)
	}
	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestAccountGetByID(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	accountRepo := NewAccountRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedAccount, err := accountRepo.Add(ctx, &account.Account{
		Name:     "Supendi Saja",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if insertedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	_, err = accountRepo.Update(ctx, &account.Account{
		ID:       insertedAccount.ID,
		Name:     "Supendi",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	fetchedAccount, err := accountRepo.GetByID(ctx, insertedAccount.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if fetchedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}

	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestAccountGetByEmail(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	accountRepo := NewAccountRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedAccount, err := accountRepo.Add(ctx, &account.Account{
		Name:     "Supendi Saja",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if insertedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	_, err = accountRepo.Update(ctx, &account.Account{
		ID:       insertedAccount.ID,
		Name:     "Supendi",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	fetchedAccount, err := accountRepo.GetByEmail(ctx, insertedAccount.Email)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedAccount == nil {
		t.Fatal("fetched account must be not nil")
	}
	if fetchedAccount.ID == "" {
		t.Fatal("fetched account id must be not an empty string")
	}
	if fetchedAccount.Email != insertedAccount.Email {
		t.Fatal("fetched account email must be equal")
	}

	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestAccountGetByPhone(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	accountRepo := NewAccountRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedAccount, err := accountRepo.Add(ctx, &account.Account{
		Name:     "Supendi Saja",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedAccount == nil {
		t.Fatal("Inserted account must be not nil")
	}
	if insertedAccount.ID == "" {
		t.Fatal("Inserted account id must be not an empty string")
	}
	_, err = accountRepo.Update(ctx, &account.Account{
		ID:       insertedAccount.ID,
		Name:     "Supendi",
		Email:    "supendi@email.com",
		Phone:    "08123",
		Password: "encrypted password",
	})
	fetchedAccount, err := accountRepo.GetByPhone(ctx, insertedAccount.Phone)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedAccount == nil {
		t.Fatal("fetched account must be not nil")
	}
	if fetchedAccount.ID == "" {
		t.Fatal("fetched account id must be not an empty string")
	}
	if fetchedAccount.Phone != insertedAccount.Phone {
		t.Fatal("fetched account phone must be equal")
	}

	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}
