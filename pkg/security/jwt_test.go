package security

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/supendi/orderan.api/pkg/identity"
)

var mySecretKey = "JambrudKatulistiwa"

type TestClaims struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

func TestGenerateAccessToken(t *testing.T) {
	handler := JWTTokenHandler{}
	newID := identity.NewID("")
	token, err := handler.GenerateAccessToken(mySecretKey, TestClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		UserID: newID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("Token should not be an empty string")
	}
}
func TestGenerateRefreshToken(t *testing.T) {
	handler := JWTTokenHandler{}
	token := handler.GenerateRefreshToken()

	if token == "" {
		t.Fatal("Token should not be an empty string")
	}
}

func TestGetClaimValue(t *testing.T) {
	handler := JWTTokenHandler{}
	newID := identity.NewID("")
	token, err := handler.GenerateAccessToken(mySecretKey, TestClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		UserID: newID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("Token should not be an empty string")
	}

	claimValue, err := handler.GetClaimValue(token, "userId", mySecretKey)
	if err != nil {
		t.Fatal(err)
	}
	if claimValue != newID {
		t.Fatalf("Claim value should be %s but got %s", newID, claimValue)
	}
}

func TestVerify(t *testing.T) {
	handler := JWTTokenHandler{}
	newID := identity.NewID("")
	token, err := handler.GenerateAccessToken(mySecretKey, TestClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		UserID: newID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("Token should not be an empty string")
	}

	verified := handler.Verify(token, mySecretKey)
	if !verified {
		t.Fatal("Verified should be true but got false")
	}
}

func TestNotVerified(t *testing.T) {
	handler := JWTTokenHandler{}
	newID := identity.NewID("")
	token, err := handler.GenerateAccessToken(mySecretKey, TestClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		UserID: newID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("Token should not be an empty string")
	}

	verified := handler.Verify(token, "wrongkey")
	if verified {
		t.Fatal("Verified should be false but got true")
	}
}
