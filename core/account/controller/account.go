package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	accountHttp "github.com/supendi/orderan.api/core/account/http"
	"github.com/supendi/orderan.api/pkg/httphelper"
)

//AccountController entry point of account service
type AccountController struct {
	responseWriter httphelper.ResponseWriter
	accountHTTP    *accountHttp.AccountHttp
	router         *chi.Mux
}

//RegisterRoutes register all account routes
func RegisterRoutes(router *chi.Mux, responseWriter httphelper.ResponseWriter, accountHTTP *accountHttp.AccountHttp) {
	controller := &AccountController{
		router:         router,
		responseWriter: responseWriter,
		accountHTTP:    accountHTTP,
	}
	controller.RegisterAccount()
}

//RegisterAccount register account route handler
func (me *AccountController) RegisterAccount() {
	me.router.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
		_, err := me.accountHTTP.RegisterAccount(r)
		me.responseWriter.Write(200, nil, err, w)
	})
}
