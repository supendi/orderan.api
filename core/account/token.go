package account

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/supendi/orderan.api/pkg/security"
)

//Claims custom claims to be encrypted
type Claims struct {
	jwt.StandardClaims
	AccountID string `json:"accountId"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

//TokenInfo represent the response model when user successfully authenticated
type TokenInfo struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

//TokenService implements TokenGenerator
type TokenService struct {
	tokenRepository TokenRepository
	tokenHandler    security.TokenHandler
	jwtKey          string
}

//NewTokenService return new TokenService instance
func NewTokenService(tokenRepository TokenRepository, tokenHandler security.TokenHandler, jwtKey string) *TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
		tokenHandler:    tokenHandler,
		jwtKey:          jwtKey,
	}
}

//GenerateTokenInfo Generates a new token info and save it into storage
func (me *TokenService) GenerateTokenInfo(ctx context.Context, account *Account) (*TokenInfo, error) {
	tokenExpireTime := time.Duration(1) * time.Hour
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireTime).Unix(),
		},
		AccountID: account.ID,
		Email:     account.Email,
		Phone:     account.Phone,
	}
	accessToken, err := me.tokenHandler.GenerateAccessToken(me.jwtKey, claims)
	if err != nil {
		return nil, err
	}

	refreshToken := me.tokenHandler.GenerateRefreshToken()

	tokenInfo := &TokenInfo{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	now := time.Now()
	_, err = me.tokenRepository.Add(ctx, &Token{
		AccessToken:    tokenInfo.AccessToken,
		RefreshToken:   tokenInfo.RefreshToken,
		RequestedCount: 0,
		Blacklisted:    false,
		CreatedAt:      now,
		ExpiredAt:      now.Add(time.Duration(120) * time.Hour), //5 days
	})

	return tokenInfo, nil
}

//GetByRefreshToken get token record from storage by refresh token
func (me *TokenService) GetByRefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	return me.tokenRepository.GetByRefreshToken(ctx, refreshToken)
}

//GetAccountID get account ID value from access token
func (me *TokenService) GetAccountID(accessToken string) (string, error) {
	return me.tokenHandler.GetClaimValue(accessToken, "accountId", me.jwtKey)
}

//Verify verifies if a token is a valid one
func (me *TokenService) Verify(accessToken string) bool {
	return me.tokenHandler.Verify(accessToken, me.jwtKey)
}
