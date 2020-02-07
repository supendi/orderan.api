package inmem

import (
	"context"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/pkg/identity"
)

//TokenRepository for working with in memory token repository
type TokenRepository struct {
	tokenList []*account.Token
}

//NewTokenRepository returns a new token repository instance
func NewTokenRepository(tokenList []*account.Token) *TokenRepository {
	return &TokenRepository{
		tokenList: tokenList,
	}
}

//Add adds new token into memory
func (me *TokenRepository) Add(ctx context.Context, token *account.Token) (*account.Token, error) {
	token.ID = identity.NewID("TOKEN_")
	me.tokenList = append(me.tokenList, token)
	return token, nil
}

//Blacklist set token blacklisted flag
func (me *TokenRepository) Blacklist(ctx context.Context, tokenID string) (*account.Token, error) {
	for _, token := range me.tokenList {
		if token.ID == tokenID {
			token.Blacklisted = true
			return token, nil
		}
	}
	return nil, nil
}

//Delete deletes existing token
func (me *TokenRepository) Delete(ctx context.Context, tokenID string) error {
	// for _, token := range me.tokenList {
	// 	if token.ID == tokenID {
	// 		token.Blacklisted = true
	// 		return token
	// 	}
	// }
	return nil
}

//GetByID get a token by ID
func (me *TokenRepository) GetByID(ctx context.Context, tokenID string) (*account.Token, error) {
	for _, token := range me.tokenList {
		if token.ID == tokenID {
			return token, nil
		}
	}
	return nil, nil
}

//GetByRefreshToken get existing token by refresh token
func (me *TokenRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (*account.Token, error) {
	for _, token := range me.tokenList {
		if token.RefreshToken == refreshToken {
			return token, nil
		}
	}
	return nil, nil
}
