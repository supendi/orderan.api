package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/database"
	"github.com/supendi/orderan.api/pkg/identity"
)

func TestTokenAdd(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	tokenRepo := NewTokenRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedToken, err := tokenRepo.Add(ctx, &account.Token{
		AccessToken:  identity.NewID("TOKEN_"),
		RefreshToken: identity.NewID(""),
		ExpiredAt:    time.Now().Add(time.Duration(1) * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if insertedToken.ID == "" {
		t.Fatal("Inserted token id must be not an empty string")
	}
	fetchedToken, err := tokenRepo.GetByID(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if fetchedToken.ID == "" {
		t.Fatal("Inserted token id must be not an empty string")
	}
	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestTokenBlacklist(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	tokenRepo := NewTokenRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedToken, err := tokenRepo.Add(ctx, &account.Token{
		AccessToken:  identity.NewID("TOKEN_"),
		RefreshToken: identity.NewID(""),
		ExpiredAt:    time.Now().Add(time.Duration(1) * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if insertedToken.ID == "" {
		t.Fatal("Inserted token id must be not an empty string")
	}
	err = tokenRepo.Blacklist(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}

	fetchedToken, err := tokenRepo.GetByID(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if !fetchedToken.Blacklisted {
		t.Fatal("Fetched token blacklist flag must me true")
	}
	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestTokenGetByID(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	tokenRepo := NewTokenRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedToken, err := tokenRepo.Add(ctx, &account.Token{
		AccessToken:  identity.NewID("TOKEN_"),
		RefreshToken: identity.NewID(""),
		ExpiredAt:    time.Now().Add(time.Duration(1) * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if insertedToken.ID == "" {
		t.Fatal("Inserted token id must be not an empty string")
	}
	err = tokenRepo.Blacklist(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}

	fetchedToken, err := tokenRepo.GetByID(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if !fetchedToken.Blacklisted {
		t.Fatal("Fetched token blacklist flag must me true")
	}
	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestTokenGetByRefreshToken(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	tokenRepo := NewTokenRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedToken, err := tokenRepo.Add(ctx, &account.Token{
		AccessToken:  identity.NewID("TOKEN_"),
		RefreshToken: identity.NewID(""),
		ExpiredAt:    time.Now().Add(time.Duration(1) * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if insertedToken.ID == "" {
		t.Fatal("Inserted token id must be not an empty string")
	}
	err = tokenRepo.Blacklist(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}

	fetchedToken, err := tokenRepo.GetByRefreshToken(ctx, insertedToken.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if !fetchedToken.Blacklisted {
		t.Fatal("Fetched token blacklist flag must me true")
	}
	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}

func TestTokenDelete(t *testing.T) {
	dbContext := database.NewDBContextTest()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	tokenRepo := NewTokenRepository(orderanDBContext.Context)
	ctx := context.Background()
	insertedToken, err := tokenRepo.Add(ctx, &account.Token{
		AccessToken:  identity.NewID("TOKEN_"),
		RefreshToken: identity.NewID(""),
		ExpiredAt:    time.Now().Add(time.Duration(1) * time.Hour),
	})
	if err != nil {
		t.Fatal(err)
	}
	if insertedToken == nil {
		t.Fatal("Inserted token must be not nil")
	}
	if insertedToken.ID == "" {
		t.Fatal("Inserted token id must be not an empty string")
	}
	err = tokenRepo.Delete(ctx, insertedToken.ID)
	if err != nil {
		t.Fatal(err)
	}

	fetchedToken, err := tokenRepo.GetByRefreshToken(ctx, insertedToken.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken != nil {
		t.Fatal("Inserted token must be nil")
	}

	err = orderanDBContext.TruncateTables()
	if err != nil {
		t.Fatal(err)
	}
	orderanDBContext.Close()
}
