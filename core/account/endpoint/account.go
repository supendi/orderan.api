package endpoint

import (
	"net/http"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/pkg/httphelper"

	"github.com/go-chi/chi"
)

//AccountController entry point and wrap account service for http layer
type AccountController struct {
	requestDecoder httphelper.RequestDecoder
	accountService *account.Service
}

//NewAccountController returns a new account http instance
func NewAccountController(requestDecoder httphelper.RequestDecoder, accountService *account.Service) *AccountController {
	return &AccountController{
		requestDecoder: requestDecoder,
		accountService: accountService,
	}
}

//RegisterAccount register new Account
func (me *AccountController) RegisterAccount(r *http.Request) (*account.Account, error) {
	var registrant account.Registrant
	err := me.requestDecoder.Decode(r, &registrant)
	if err != nil {
		return nil, err
	}
	registeredAccount, err := me.accountService.RegisterAccount(r.Context(), &registrant)
	return registeredAccount, err
}

//GetAccount gets an existing Account
func (me *AccountController) GetAccount(r *http.Request) (*account.Account, error) {
	getRequest := &account.GetRequest{}
	getRequest.AccountID = me.requestDecoder.URLParam(r, "accountId")
	return me.accountService.GetAccount(r.Context(), getRequest)
}

//UpdateAccount updates an existing account
func (me *AccountController) UpdateAccount(r *http.Request) (*account.Account, error) {
	var updateRequest account.UpdateRequest
	err := me.requestDecoder.Decode(r, &updateRequest)
	if err != nil {
		return nil, err
	}
	updateRequest.AccountID = me.requestDecoder.URLParam(r, "accountId")
	return me.accountService.UpdateAccount(r.Context(), &updateRequest)
}

//RegisterRoutes register all account routes
func RegisterRoutes(router *chi.Mux, responseWriter httphelper.ResponseWriter, accountCtrl *AccountController) {
	router.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
		accountInfo, err := accountCtrl.RegisterAccount(r)
		responseWriter.Write(200, accountInfo, err, w)
	})
	router.Get("/accounts/{accountId}", func(w http.ResponseWriter, r *http.Request) {
		accountInfo, err := accountCtrl.GetAccount(r)
		responseWriter.Write(200, accountInfo, err, w)
	})
	router.Put("/accounts/{accountId}", func(w http.ResponseWriter, r *http.Request) {
		accountInfo, err := accountCtrl.UpdateAccount(r)
		responseWriter.Write(200, accountInfo, err, w)
	})
}
