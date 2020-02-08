package endpoint

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/pkg/httphelper"
	"github.com/supendi/orderan.api/pkg/validator"
)

//AccountController entry point and wrap account service for http layer
type AccountController struct {
	decoder        httphelper.RequestDecoder
	validator      validator.Validator
	accountService *account.Service
}

//NewAccountController returns a new account controller instance. Its only wrapping account service, providing service method from decoded http request to required param type
func NewAccountController(decoder httphelper.RequestDecoder, validator validator.Validator, accountService *account.Service) *AccountController {
	return &AccountController{
		decoder:        decoder,
		validator:      validator,
		accountService: accountService,
	}
}

//RegisterAccount register new Account
func (me *AccountController) RegisterAccount(r *http.Request) (*account.Account, error) {
	var registrant account.Registrant
	err := me.decoder.DecodeAndValidate(r, &registrant)
	if err != nil {
		return nil, err
	}
	registeredAccount, err := me.accountService.RegisterAccount(r.Context(), &registrant)
	return registeredAccount, err
}

//GetAccount gets an existing Account
func (me *AccountController) GetAccount(r *http.Request) (*account.Account, error) {
	getRequest := &account.GetRequest{}
	getRequest.AccountID = me.decoder.URLParam(r, "accountId")
	err := me.validator.Validate(getRequest)
	if err != nil {
		return nil, err
	}
	return me.accountService.GetAccount(r.Context(), getRequest)
}

//UpdateAccount updates an existing account
func (me *AccountController) UpdateAccount(r *http.Request) (*account.Account, error) {
	var updateRequest account.UpdateRequest
	err := me.decoder.Decode(r, &updateRequest)
	if err != nil {
		return nil, err
	}
	updateRequest.AccountID = me.decoder.URLParam(r, "accountId")
	err = me.validator.Validate(&updateRequest)
	if err != nil {
		return nil, err
	}
	return me.accountService.UpdateAccount(r.Context(), &updateRequest)
}

func mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("asd")
		//Decrypt jwt token here
		next.ServeHTTP(w, r)
	})
}

//RegisterRoutes register all account routes
func RegisterRoutes(router *chi.Mux, responseWriter httphelper.ResponseWriter, accountCtrl *AccountController) {

	router.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
		accountInfo, err := accountCtrl.RegisterAccount(r)
		responseWriter.Write(200, accountInfo, err, w)
	})
	router.Group(func(r chi.Router) {
		r.Use(mw)
		r.Get("/accounts/{accountId}", func(w http.ResponseWriter, request *http.Request) {
			accountInfo, err := accountCtrl.GetAccount(request)
			responseWriter.Write(200, accountInfo, err, w)
		})
		r.Put("/accounts/{accountId}", func(w http.ResponseWriter, request *http.Request) {
			accountInfo, err := accountCtrl.UpdateAccount(request)
			responseWriter.Write(200, accountInfo, err, w)
		})
	})
}
