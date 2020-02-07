package account

import (
	"testing"

	"github.com/supendi/orderan.api/pkg/identity"
	"github.com/supendi/orderan.api/pkg/security"
)

func TestGenerateTokenInfo(t *testing.T) {
	newAccountID := identity.NewID("A_")
	service := NewTokenService(&security.JWTTokenHandler{})
	tokenInfo, err := service.GenerateTokenInfo(&Account{
		ID:    newAccountID,
		Email: "irpan@gmail.com",
		Phone: "0813",
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

func TestGetAccountID(t *testing.T) {
	newAccountID := identity.NewID("A_")
	service := NewTokenService(&security.JWTTokenHandler{})
	tokenInfo, err := service.GenerateTokenInfo(&Account{
		ID:    newAccountID,
		Email: "irpan@gmail.com",
		Phone: "0813",
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
	accountIDFromToken, err := service.GetAccountID(tokenInfo.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if accountIDFromToken != newAccountID {
		t.Fatal("Invalid account ID")
	}
}

func TestVerify(t *testing.T) {
	newAccountID := identity.NewID("A_")
	service := NewTokenService(&security.JWTTokenHandler{})
	tokenInfo, err := service.GenerateTokenInfo(&Account{
		ID:    newAccountID,
		Email: "irpan@gmail.com",
		Phone: "0813",
	})
	if err != nil {
		t.Fatal(err)
	}
	if tokenInfo == nil {
		t.Fatal("tokenInfo shouldnt be nil")
	}
	isValidToken := service.Verify(tokenInfo.AccessToken)
	if !isValidToken {
		t.Fatal("Invalid token")
	}
}
