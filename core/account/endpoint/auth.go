package endpoint

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/core/account"
	"github.com/supendi/orderan.api/pkg/httphelper"
	"github.com/supendi/orderan.api/pkg/validator"
)

//AuthController is an auth service wrapper
type AuthController struct {
	decoder     httphelper.RequestDecoder
	validator   validator.Validator
	authService *account.AuthService
}

//NewAuthController returns a new account controller instance. Its only wrapping account service, providing service method from decoded http request to required param type
func NewAuthController(decoder httphelper.RequestDecoder, validator validator.Validator, authService *account.AuthService) *AuthController {
	return &AuthController{
		decoder:     decoder,
		validator:   validator,
		authService: authService,
	}
}

//Authenticate Authenticates user
func (me *AuthController) Authenticate(r *http.Request) (*account.TokenInfo, error) {
	var loginRequest account.LoginRequest
	err := me.decoder.DecodeBodyAndValidate(r, &loginRequest)
	if err != nil {
		return nil, err
	}
	return me.authService.Authenticate(r.Context(), &loginRequest)
}

//RenewAccessToken renew access token
func (me *AuthController) RenewAccessToken(r *http.Request) (*account.TokenInfo, error) {
	var req account.RenewTokenRequest
	err := me.decoder.DecodeBodyAndValidate(r, &req)
	if err != nil {
		return nil, err
	}
	return me.authService.RenewAccessToken(r.Context(), &req)
}

//RegisterAuthRoutes register all auth routes
func RegisterAuthRoutes(router *chi.Mux, responseWriter httphelper.ResponseWriter, authCtrl *AuthController) {
	router.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		tokenInfo, err := authCtrl.Authenticate(r)
		responseWriter.Write(200, tokenInfo, err, w)
	})
	router.Post("/refresh-token", func(w http.ResponseWriter, r *http.Request) {
		tokenInfo, err := authCtrl.RenewAccessToken(r)
		responseWriter.Write(200, tokenInfo, err, w)
	})
}
