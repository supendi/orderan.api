package http

import (
	"net/http"

	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/pkg/httphelper"
)

type AccountHttp struct {
	requestDecoder httphelper.RequestDecoder
	accountService *account.Service
}

func NewAccountHTTP(requestDecoder httphelper.RequestDecoder, accountService *account.Service) *AccountHttp {
	return &AccountHttp{
		requestDecoder: requestDecoder,
		accountService: accountService,
	}
}

//Register register new Account
func (me *AccountHttp) RegisterAccount(r *http.Request) (*account.Account, error) {
	var registrant account.Registrant
	err := me.requestDecoder.Decode(r, &registrant)
	if err != nil {
		return nil, err
	}
	return me.accountService.RegisterAccount(r.Context(), &registrant)
}
