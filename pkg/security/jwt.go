package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/supendi/orderan.api/pkg/identity"
)

//TokenHandler specifies the functionalties contract of token handler
type TokenHandler interface {
	GenerateAccessToken(signKey string, claims jwt.Claims) (string, error)
	GenerateRefreshToken() string
	GetClaimValue(accessToken string, claimKey string, jwtKey string) (string, error)
	Verify(accessToken string, jwtKey string) bool
}

//JWTTokenHandler implement token handler
type JWTTokenHandler struct{}

//GenerateAccessToken generate a signed token string
func (me *JWTTokenHandler) GenerateAccessToken(signKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey := []byte(signKey)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//GenerateRefreshToken generates a refresh token
func (me *JWTTokenHandler) GenerateRefreshToken() string {
	return identity.NewID("")
}

//GetClaimValue get claim value from access token
func (me *JWTTokenHandler) GetClaimValue(accessToken string, claimKey string, jwtKey string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", err
	}
	claimValue := token.Claims.(jwt.MapClaims)[claimKey]
	if claimValue != nil {
		return claimValue.(string), nil
	}
	return "", nil
}

//Verify verifies that a JWT is valid
func (me *JWTTokenHandler) Verify(accessToken string, jwtKey string) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
