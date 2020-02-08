package account_test

import (
	"testing"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/account/inmem"
	"github.com/supendi/orderan.api/pkg/security"
)

func TestAuthenticate(t *testing.T) {
	//Preparing registrant and create a new account
	registrant := &account.Registrant{
		Name:     "ressy",
		Email:    "ressy@gmail.com",
		Phone:    "08144",
		Password: "passwordnyaKuat",
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

	//Test authentication
	tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenService := account.NewTokenService(tokenRepo, &security.JWTTokenHandler{}, "SECRETKEY")
	authService := account.NewAuthService(tokenService, GetAccountRepo(), account.NewBCryptHasher())
	tokenInfo, err := authService.Authenticate(newContext, &account.LoginRequest{
		Username: registrant.Email,
		Password: registrant.Password,
	})
	if err != nil {
		t.Fatal(err)
	}
	if tokenInfo == nil {
		t.Fatal("tokenInfo shouldnt be nil")
	}
	if tokenInfo.AccessToken == "" {
		t.Fatal("Access token shouldnt be an empty string")
	}
	if tokenInfo.RefreshToken == "" {
		t.Fatal("Refresh token shouldnt be an empty string")
	}
}

func TestAuthenticateFailed(t *testing.T) {
	//Preparing registrant and create a new account
	registrant := &account.Registrant{
		Name:     "lele",
		Email:    "lele@gmail.com",
		Phone:    "081422",
		Password: "passwordnyaKuat",
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

	//Test authentication
	tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenService := account.NewTokenService(tokenRepo, &security.JWTTokenHandler{}, "SECRETKEY")
	authService := account.NewAuthService(tokenService, GetAccountRepo(), account.NewBCryptHasher())
	tokenInfo, err := authService.Authenticate(newContext, &account.LoginRequest{
		Username: registrant.Email,
		Password: "wrong password",
	})
	if err == nil {
		t.Fatal("Should return auth error")
	}
	if tokenInfo != nil {
		t.Fatal("tokenInfo should be nil")
	}
}

func TestRenewAccessToken(t *testing.T) {
	//Preparing registrant and create a new account
	registrant := &account.Registrant{
		Name:     "suanto",
		Email:    "suanto@gmail.com",
		Phone:    "08199",
		Password: "passwordnyaKuat",
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

	//Test authentication
	tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenService := account.NewTokenService(tokenRepo, &security.JWTTokenHandler{}, "SECRETKEY")
	authService := account.NewAuthService(tokenService, GetAccountRepo(), account.NewBCryptHasher())
	tokenInfo, err := authService.Authenticate(newContext, &account.LoginRequest{
		Username: registrant.Email,
		Password: registrant.Password,
	})
	if err != nil {
		t.Fatal(err)
	}
	if tokenInfo == nil {
		t.Fatal("tokenInfo shouldnt be nil")
	}
	if tokenInfo.AccessToken == "" {
		t.Fatal("Access token shouldnt be an empty string")
	}
	if tokenInfo.RefreshToken == "" {
		t.Fatal("Refresh token shouldnt be an empty string")
	}

	newAccessToken, err := authService.RenewAccessToken(newContext, &account.RenewTokenRequest{
		AccessToken:  tokenInfo.AccessToken,
		RefreshToken: tokenInfo.RefreshToken,
	})
	if err != nil {
		t.Fatal(err)
	}

	if newAccessToken == nil {
		t.Fatal("newAccessToken shouldnt be nil")
	}
	if newAccessToken.AccessToken == "" {
		t.Fatal("Access token shouldnt be an empty string")
	}
	if newAccessToken.RefreshToken == "" {
		t.Fatal("Refresh token shouldnt be an empty string")
	}

	tokenRecord, err := tokenRepo.GetByRefreshToken(newContext, newAccessToken.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	if tokenRecord == nil {
		t.Fatal("token shouldnt be nil")
	}
}
