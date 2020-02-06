package account

import (
	"context"
	"time"
)

type (
	//ResfreshToken represent refresh token
	ResfreshToken struct {
		ID             string
		AccessToken    string
		RefreshToken   string
		RequestedCount int
		Blacklisted    bool
		CreatedAt      time.Time
		UpdatedAt      *time.Time
	}
)

type (
	//LoginRequest represent the model of login request
	LoginRequest struct {
		Username string
		Password string
	}

	//RefreshTokenRequest represent the model while requesting a new access token
	RefreshTokenRequest struct {
		AccessToken  string
		RefreshToken string
	}

	//TokenInfo represent the response model when user successfully authenticated
	TokenInfo struct {
		AccessToken  string
		RefreshToken string
	}
)

type (
	//TokenRepository specifies the functionalities for working with token data storage
	TokenRepository interface {
		Add(ctx context.Context, token *Token) (*Token, error)
		Blacklist(ctx context.Context, tokenID string) (*Token, error)
		Delete(ctx context.Context, tokenID string) error
		GetByID(ctx context.Context, tokenID string) (*Token, error)
		GetByRefreshToken(ctx context.Context, refreshToken string) (*Token, error)
	}

	//SecurityTokenHandler specifies the functionalties contract of token handler
	SecurityTokenHandler interface {
		GenerateToken(account *Account) string
	}
)

//AuthService specifies the functionalies user authentication
type AuthService struct {
	accountRepository Repository
	tokenRepository   TokenRepository
	passwordHasher    PasswordHasher
}

//Authenticate authenticates user
func (me *AuthService) Authenticate(req *LoginRequest) {
	existingAccount := me.accountRepository
}
