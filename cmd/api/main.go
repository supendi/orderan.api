package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	accountController "github.com/supendi/orderan.api/core/account/controller"
	accountHttp "github.com/supendi/orderan.api/core/account/http"
	"github.com/supendi/orderan.api/core/account/inmem"

	"github.com/supendi/orderan.api/pkg/httphelper"
)

func main() {
	r := chi.NewRouter()
	hasher := account.NewBCryptHasher()
	accountRepo := inmem.NewAccountRepository([]*account.Account{})
	accountService := account.NewAccountService(accountRepo, hasher)
	accountHTTP := accountHttp.NewAccountHTTP(&httphelper.RequestHandler{}, accountService)
	accountController.RegisterRoutes(r, &httphelper.ResponseHandler{}, accountHTTP)

	http.ListenAndServe(":8080", r)
}
