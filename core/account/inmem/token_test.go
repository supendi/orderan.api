package inmem

import (
	"context"
	"testing"
	"time"

	"github.com/supendi/orderan.api/core/account"
)

func TestAddToken(t *testing.T) {
	token := &account.Token{
		AccessToken:  "asfd",
		RefreshToken: "fdsa",
		Blacklisted:  false,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Duration(2) * time.Hour),
	}
	tokenRepo := NewTokenRepository([]*account.Token{})
	newContext := context.Background()
	_, err := tokenRepo.Add(newContext, token)
	if err != nil {
		t.Fatal(err)
	}
	fetchedToken, err := tokenRepo.GetByID(newContext, token.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("fetchedToken shouldnt be nil")
	}
}

func TestBlacklistToken(t *testing.T) {
	token := &account.Token{
		AccessToken:  "asfd",
		RefreshToken: "fdsa",
		Blacklisted:  false,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Duration(2) * time.Hour),
	}
	tokenRepo := NewTokenRepository([]*account.Token{})
	newContext := context.Background()
	_, err := tokenRepo.Add(newContext, token)
	if err != nil {
		t.Fatal(err)
	}
	fetchedToken, err := tokenRepo.Blacklist(newContext, token.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("fetchedToken shouldnt be nil")
	}

	if !fetchedToken.Blacklisted {
		t.Fatal("fetchedToken should be blacklisted")
	}
}

func TestDeleteToken(t *testing.T) {
	token := &account.Token{
		AccessToken:  "asfd",
		RefreshToken: "fdsa",
		Blacklisted:  false,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Duration(2) * time.Hour),
	}
	tokenRepo := NewTokenRepository([]*account.Token{})
	newContext := context.Background()
	_, err := tokenRepo.Add(newContext, token)
	if err != nil {
		t.Fatal(err)
	}
	err = tokenRepo.Delete(newContext, token.ID)
	if err != nil {
		t.Fatal(err)
	}
	fetchedToken, err := tokenRepo.GetByID(newContext, token.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken != nil {
		t.Fatal("fetchedToken should be nil, as a prove it's been deleted")
	}

}

func TestGetByID(t *testing.T) {
	token := &account.Token{
		AccessToken:  "asfd",
		RefreshToken: "fdsa",
		Blacklisted:  false,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Duration(2) * time.Hour),
	}
	tokenRepo := NewTokenRepository([]*account.Token{})
	newContext := context.Background()
	_, err := tokenRepo.Add(newContext, token)
	if err != nil {
		t.Fatal(err)
	}
	fetchedToken, err := tokenRepo.GetByID(newContext, token.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("fetchedToken shouldnt be nil")
	}
}

func TestGetByRefreshToken(t *testing.T) {
	token := &account.Token{
		AccessToken:  "asfd",
		RefreshToken: "fdsa",
		Blacklisted:  false,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Duration(2) * time.Hour),
	}
	tokenRepo := NewTokenRepository([]*account.Token{})
	newContext := context.Background()
	_, err := tokenRepo.Add(newContext, token)
	if err != nil {
		t.Fatal(err)
	}
	fetchedToken, err := tokenRepo.GetByRefreshToken(newContext, token.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedToken == nil {
		t.Fatal("fetchedToken shouldnt be nil")
	}
}
