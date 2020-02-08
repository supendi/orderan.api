package main

import (
	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	accountEndpoint "github.com/supendi/orderan.api/core/account/endpoint"
	"github.com/supendi/orderan.api/core/account/inmem"

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

	//Account
	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountService := account.NewAccountService(accountRepo, hasher)
	accountController := accountEndpoint.NewAccountController(requestHandler, modelValidator, accountService)
	accountEndpoint.RegisterAccountRoutes(router, responseHandler, tokenHandler, jwtKey, accountController)

	//Auth
	tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenService := account.NewTokenService(tokenRepo, tokenHandler, jwtKey)
	authService := account.NewAuthService(tokenService, accountRepo, hasher)
	authController := accountEndpoint.NewAuthController(requestHandler, modelValidator, authService)
	accountEndpoint.RegisterAuthRoutes(router, responseHandler, authController)
}
