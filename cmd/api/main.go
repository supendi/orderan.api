package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	accountEndpoint "github.com/supendi/orderan.api/core/account/endpoint"
	"github.com/supendi/orderan.api/core/account/inmem"

	"github.com/supendi/orderan.api/pkg/httphelper"
)

func main() {
	r := chi.NewRouter()
	hasher := account.NewBCryptHasher()
	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountService := account.NewAccountService(accountRepo, hasher)
	accountController := accountEndpoint.NewAccountController(&httphelper.RequestHandler{}, accountService)
	accountEndpoint.RegisterRoutes(r, &httphelper.ResponseHandler{}, accountController)

	http.ListenAndServe(":8080", r)
}
