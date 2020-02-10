package account

import (
	"context"
	"time"

	"github.com/supendi/orderan.api/pkg/errors"
)

//AuthError returns app err with invalid username or password message
func AuthError() error {
	appErr := errors.NewAppError("Invalid username or password.")
	return appErr
}

//InvalidTokenError returns app err with invalid token message
func InvalidTokenError() error {
	appErr := errors.NewAppError("Invalid access token or refresh token.")
	return appErr
}

//ExpiredTokenError returns app err with expired token message
func ExpiredTokenError() error {
	appErr := errors.NewAppError("Refresh token is expired.")
	return appErr
}

type (
	//Token represent token model
	Token struct {
		ID             string
		AccessToken    string
		RefreshToken   string
		RequestedCount int
		Blacklisted    bool
		ExpiredAt      time.Time
		CreatedAt      time.Time
		UpdatedAt      *time.Time
	}
)

type (
	//LoginRequest represent the model of login request
	LoginRequest struct {
		Username string `validate:"required"`
		Password string `validate:"required"`
	}

	//RenewTokenRequest represent the model while requesting a new access token
	RenewTokenRequest struct {
		AccessToken  string `validate:"required"`
		RefreshToken string `validate:"required"`
	}
)

type (
	//TokenRepository specifies the functionalities for working with token data storage
	TokenRepository interface {
		Add(ctx context.Context, token *Token) (*Token, error)
		Blacklist(ctx context.Context, tokenID string) error
		Delete(ctx context.Context, tokenID string) error
		GetByID(ctx context.Context, tokenID string) (*Token, error)
		GetByRefreshToken(ctx context.Context, refreshToken string) (*Token, error)
	}
)

//AuthService specifies the functionalies user authentication
type AuthService struct {
	tokenService      *TokenService
	accountRepository Repository
	passwordHasher    PasswordHasher
}

//NewAuthService returns a new account service
func NewAuthService(tokenService *TokenService, accountRepository Repository, passwordHasher PasswordHasher) *AuthService {
	return &AuthService{
		tokenService:      tokenService,
		accountRepository: accountRepository,
		passwordHasher:    passwordHasher,
	}
}

//Authenticate authenticates user
func (me *AuthService) Authenticate(ctx context.Context, req *LoginRequest) (*TokenInfo, error) {
	//Find account by email first
	existingAccount, err := me.accountRepository.GetByEmail(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	//If not found, find by phone number
	if existingAccount == nil {
		existingAccount, err = me.accountRepository.GetByPhone(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		if existingAccount == nil {
			return nil, AuthError()
		}
	}

	passwordIsValid, err := me.passwordHasher.Verify(req.Password, existingAccount.Password)
	if err != nil {
		return nil, err
	}
	if !passwordIsValid {
		return nil, AuthError()
	}

	tokenInfo, err := me.tokenService.GenerateTokenInfo(ctx, existingAccount)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

//RenewAccessToken renew access token by providing its access token and its refresh token
func (me *AuthService) RenewAccessToken(ctx context.Context, req *RenewTokenRequest) (*TokenInfo, error) {
	existingToken, err := me.tokenService.GetByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if existingToken == nil {
		return nil, InvalidTokenError()
	}
	//Access token must be verified too, because it is saved together with refresh token in data storage
	if existingToken.AccessToken != req.AccessToken {
		return nil, InvalidTokenError()
	}

	tokenIsExpired := time.Now().After(existingToken.ExpiredAt)
	if tokenIsExpired {
		return nil, ExpiredTokenError()
	}

	accountID, err := me.tokenService.GetAccountID(req.AccessToken)
	if err != nil {
		return nil, err
	}
	existingAccount, err := me.accountRepository.GetByID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	if existingAccount == nil {
		return nil, AccountNotFoundError(accountID)
	}

	tokenInfo, err := me.tokenService.GenerateTokenInfo(ctx, existingAccount)

	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
