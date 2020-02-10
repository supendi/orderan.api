package main

import (
	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	accountEndpoint "github.com/supendi/orderan.api/core/account/endpoint"

	//accountInmem "github.com/supendi/orderan.api/core/account/inmem"
	accountPostgres "github.com/supendi/orderan.api/core/account/postgres"
	"github.com/supendi/orderan.api/core/database"

	"github.com/supendi/orderan.api/pkg/httphelper"

	"github.com/supendi/orderan.api/pkg/security"
	"github.com/supendi/orderan.api/pkg/validator"
)

func run(router *chi.Mux, jwtKey string) {
	requestHandler := &httphelper.RequestHandler{}
	responseHandler := &httphelper.ResponseHandler{}
	modelValidator := &validator.ModelValidator{}
	tokenHandler := &security.JWTTokenHandler{}
	hasher := account.NewBCryptHasher()

	dbContext := database.NewDBContext()
	orderanDBContext := database.NewOrderanDBContext(dbContext)

	//Account
	//accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountRepo := accountPostgres.NewAccountRepository(orderanDBContext.Context)
	accountService := account.NewAccountService(accountRepo, hasher)
	accountController := accountEndpoint.NewAccountController(requestHandler, modelValidator, accountService)
	accountEndpoint.RegisterAccountRoutes(router, responseHandler, tokenHandler, jwtKey, accountController)

	//Auth
	//tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenRepo := accountPostgres.NewTokenRepository(orderanDBContext.Context)
	tokenService := account.NewTokenService(tokenRepo, tokenHandler, jwtKey)
	authService := account.NewAuthService(tokenService, accountRepo, hasher)
	authController := accountEndpoint.NewAuthController(requestHandler, modelValidator, authService)
	accountEndpoint.RegisterAuthRoutes(router, responseHandler, authController)
}
