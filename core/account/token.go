package account

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/supendi/orderan.api/pkg/security"
)

var jwtKey string = ""

func getJWTkey() string {
	if jwtKey == "" {
		jwtKey = os.Getenv("JWT_KEY")
		if jwtKey == "" {
			jwtKey = "PengenTinggalDiBandungBrooo"
		}
	}
	return jwtKey
}

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
}

//NewTokenService return new TokenService instance
func NewTokenService(tokenRepository TokenRepository, tokenHandler security.TokenHandler) *TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
		tokenHandler:    tokenHandler,
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
	accessToken, err := me.tokenHandler.GenerateAccessToken(getJWTkey(), claims)
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

//GetAccountID Generates a new token info
func (me *TokenService) GetAccountID(accessToken string) (string, error) {
	return me.tokenHandler.GetClaimValue(accessToken, "accountId", getJWTkey())
}

//Verify verifies if a token is a valid one
func (me *TokenService) Verify(accessToken string) bool {
	return me.tokenHandler.Verify(accessToken, getJWTkey())
}
