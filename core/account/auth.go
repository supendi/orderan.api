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
	appErr := errors.NewAppError("Refresh token is already expired.")
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
		Username string
		Password string
	}

	//RenewTokenRequest represent the model while requesting a new access token
	RenewTokenRequest struct {
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
		GenerateAccessToken(account *Account) (string, error)
		GenerateRefreshToken() (string, error)
		GetSubValue(accessToken string) (string, error)
	}
)

//AuthService specifies the functionalies user authentication
type AuthService struct {
	accountRepository Repository
	tokenRepository   TokenRepository
	passwordHasher    PasswordHasher
	tokenHandler      SecurityTokenHandler
}

//generateTokenInfo Generates a new token info
func (me *AuthService) generateTokenInfo(ctx context.Context, account *Account) (*TokenInfo, error) {
	refreshToken, err := me.tokenHandler.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	accessToken, err := me.tokenHandler.GenerateAccessToken(account)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	_, err = me.tokenRepository.Add(ctx, &Token{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		RequestedCount: 0,
		Blacklisted:    false,
		CreatedAt:      now,
		ExpiredAt:      now.Add(time.Duration(120) * time.Hour), //5 days
	})
	if err != nil {
		return nil, err
	}

	tokenInfo := &TokenInfo{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenInfo, nil
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
	tokenInfo, err := me.generateTokenInfo(ctx, existingAccount)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

//RenewAccessToken renew access token by profiding its access token and its refresh token
func (me *AuthService) RenewAccessToken(ctx context.Context, req *RenewTokenRequest) (*TokenInfo, error) {
	existingToken, err := me.tokenRepository.GetByRefreshToken(ctx, req.RefreshToken)
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

	tokenIsExpired := existingToken.ExpiredAt.After(time.Now())
	if tokenIsExpired {
		return nil, ExpiredTokenError()
	}

	accountID, err := me.tokenHandler.GetSubValue(req.AccessToken)
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

	tokenInfo, err := me.generateTokenInfo(ctx, existingAccount)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
