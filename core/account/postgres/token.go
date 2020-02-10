package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/supendi/dbx"
	"github.com/supendi/orderan.api/core/account"
)

//TokenRepository implements account.Repository
type TokenRepository struct {
	db *dbx.Context
}

//NewTokenRepository returns a new Token Repository instance
func NewTokenRepository(dbContext *dbx.Context) *TokenRepository {
	return &TokenRepository{
		db: dbContext,
	}
}

//Add a new token into storage
func (me *TokenRepository) Add(ctx context.Context, newToken *account.Token) (*account.Token, error) {
	newToken.ID = uuid.New().String()

	statement := dbx.NewStatement("INSERT INTO token(id, access_token, refresh_token, requested_count, blacklisted, expired_at, created_at) VALUES (:id, :access_token, :refresh_token, :requested_count, :blacklisted, :expired_at, NOW())")
	statement.AddParameter("id", newToken.ID)
	statement.AddParameter("access_token", newToken.AccessToken)
	statement.AddParameter("refresh_token", newToken.RefreshToken)
	statement.AddParameter("requested_count", newToken.RequestedCount)
	statement.AddParameter("blacklisted", newToken.Blacklisted)
	statement.AddParameter("expired_at", newToken.ExpiredAt)

	me.db.AddStatement(statement)

	_, err := me.db.SaveChanges(ctx)

	if err != nil {
		return nil, err
	}

	return newToken, nil
}

//Blacklist updates an existing token blacklisted flag to true
func (me *TokenRepository) Blacklist(ctx context.Context, tokenID string) error {
	statement := dbx.NewStatement("UPDATE token SET blacklisted=true, updated_at = NOW() WHERE id = :id")
	statement.AddParameter("id", tokenID)

	_, err := me.db.SaveChanges(ctx)

	if err != nil {
		return err
	}

	return nil
}

//GetByID returns an token by ID
func (me *TokenRepository) GetByID(ctx context.Context, id string) (*account.Token, error) {
	statement := dbx.NewStatement("SELECT * FROM token WHERE id = :id")
	statement.AddParameter("id", id)

	rows, err := me.db.QueryStatementContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var retrievedToken *account.Token

	for rows.Next() {
		retrievedToken = &account.Token{}
		err := rows.Scan(&retrievedToken.ID, &retrievedToken.AccessToken, &retrievedToken.RefreshToken, &retrievedToken.RequestedCount, &retrievedToken.Blacklisted, &retrievedToken.ExpiredAt, &retrievedToken.CreatedAt, &retrievedToken.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return retrievedToken, nil
	}

	return nil, nil
}

//GetByRefreshToken returns an token by email
func (me *TokenRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (*account.Token, error) {
	statement := dbx.NewStatement("SELECT * FROM token WHERE refresh_token = :refresh_token")
	statement.AddParameter("refresh_token", refreshToken)

	rows, err := me.db.QueryStatementContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	var retrievedToken *account.Token

	for rows.Next() {
		retrievedToken = &account.Token{}
		err := rows.Scan(&retrievedToken.ID, &retrievedToken.AccessToken, &retrievedToken.RefreshToken, &retrievedToken.RequestedCount, &retrievedToken.Blacklisted, &retrievedToken.ExpiredAt, &retrievedToken.CreatedAt, &retrievedToken.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return retrievedToken, nil
	}

	return nil, nil
}
