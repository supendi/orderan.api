package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/core/account/controller"
	"github.com/supendi/orderan.api/core/account/inmem"

	"github.com/supendi/orderan.api/pkg/httphelper"
)

func main() {
	r := chi.NewRouter()
	hasher := account.NewBCryptHasher()
	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountService := account.NewAccountService(accountRepo, hasher)
	accountController := controller.NewAccountController(&httphelper.RequestHandler{}, accountService)
	controller.RegisterRoutes(r, &httphelper.ResponseHandler{}, accountController)

	http.ListenAndServe(":8080", r)
}
