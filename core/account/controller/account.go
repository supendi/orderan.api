package controller

import (
	"net/http"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/pkg/httphelper"

	"github.com/go-chi/chi"
)

//AccountController wrap account service for http layer
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

//AccountRoute entry point of account service
type AccountRoute struct {
	router         *chi.Mux
	responseWriter httphelper.ResponseWriter
	accountCtrl    *AccountController
}

//RegisterRoutes register all account routes
func RegisterRoutes(router *chi.Mux, responseWriter httphelper.ResponseWriter, accountCtrl *AccountController) {
	route := &AccountRoute{
		router:         router,
		responseWriter: responseWriter,
		accountCtrl:    accountCtrl,
	}
	route.RegisterAccount()
}

//RegisterAccount register account route handler
func (me *AccountRoute) RegisterAccount() {
	me.router.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
		accountInfo, err := me.accountCtrl.RegisterAccount(r)
		me.responseWriter.Write(200, accountInfo, err, w)
	})
}
