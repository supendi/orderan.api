package account

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/supendi/orderan.api/pkg/security"
)

var jwtKey string

func getJWTkey() string {
	if jwtKey == "" {
		jwtKey = os.Getenv("JWT_KEY")
		if jwtKey == "" {
			jwtKey = "PengenTinggalDiBandungBrooo"
		}
	}
	return jwtKey
}

//Claims custome claims to be encrypted
type Claims struct {
	jwt.StandardClaims
	AccountID string `json:accountId`
	Email     string `json:"username"`
	Phone     string `json:"phone"`
}

//TokenInfo represent the response model when user successfully authenticated
type TokenInfo struct {
	AccessToken  string
	RefreshToken string
}

//TokenService implements TokenGenerator
type TokenService struct {
	tokenHandler security.TokenHandler
}

//NewTokenService return new TokenService instance
func NewTokenService(tokenHandler security.TokenHandler) *TokenService {
	return &TokenService{
		tokenHandler: tokenHandler,
	}
}

//GenerateTokenInfo Generates a new token info
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

	return tokenInfo, nil
}

//GetAccountID Generates a new token info
func (me *TokenService) GetAccountID(accessToken string) (string, error) {
	return me.tokenHandler.GetClaimValue(accessToken, "accountId", getJWTkey())
}
