package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	accountEndpoint "github.com/supendi/orderan.api/core/account/endpoint"
	"github.com/supendi/orderan.api/core/account/inmem"

	"github.com/supendi/orderan.api/pkg/httphelper"

	"github.com/supendi/orderan.api/pkg/security"
	"github.com/supendi/orderan.api/pkg/validator"
)

func main() {
	jwtKey := "PengenTinggalDiBandungBrooo"
	requestHandler := &httphelper.RequestHandler{}
	responseHandler := &httphelper.ResponseHandler{}
	modelValidator := &validator.ModelValidator{}
	tokenHandler := &security.JWTTokenHandler{}
	r := chi.NewRouter()
	hasher := account.NewBCryptHasher()

	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountService := account.NewAccountService(accountRepo, hasher)
	accountController := accountEndpoint.NewAccountController(requestHandler, modelValidator, accountService)
	accountEndpoint.RegisterAccountRoutes(r, responseHandler, tokenHandler, jwtKey, accountController)

	tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenService := account.NewTokenService(tokenRepo, tokenHandler)
	authService := account.NewAuthService(tokenService, accountRepo, hasher)
	authController := accountEndpoint.NewAuthController(requestHandler, modelValidator, authService)
	accountEndpoint.RegisterAuthRoutes(r, responseHandler, authController)

	http.ListenAndServe(":8080", r)
}
