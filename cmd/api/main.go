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
	r := chi.NewRouter()
	hasher := account.NewBCryptHasher()
	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountService := account.NewAccountService(accountRepo, hasher)
	accountController := accountEndpoint.NewAccountController(&httphelper.RequestHandler{}, &validator.ModelValidator{}, accountService)
	accountEndpoint.RegisterAccountRoutes(r, &httphelper.ResponseHandler{}, accountController)

	tokenRepo := inmem.NewTokenRepository([]*account.Token{})
	tokenService := account.NewTokenService(tokenRepo, &security.JWTTokenHandler{})
	authService := account.NewAuthService(tokenService, accountRepo, hasher)
	authController := accountEndpoint.NewAuthController(&httphelper.RequestHandler{}, &validator.ModelValidator{}, authService)
	accountEndpoint.RegisterAuthRoutes(r, &httphelper.ResponseHandler{}, authController)

	http.ListenAndServe(":8080", r)
}
