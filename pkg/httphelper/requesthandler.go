package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/pkg/validator"
)

//RequestDecoder decode http request to a spesific model
type RequestDecoder interface {
	Decode(r *http.Request, decodeTo interface{}) error
	DecodeAndValidate(r *http.Request, model interface{}) error
	URLParam(r *http.Request, key string) string
}

//RequestHandler implements RequestDecoder
type RequestHandler struct {
}

//Decode decode request body into specified param
func (me *RequestHandler) Decode(r *http.Request, decodeTo interface{}) error {
	err := json.NewDecoder(r.Body).Decode(decodeTo)
	return err
}

//DecodeAndValidate decode the request body, and then validates the struct
func (me *RequestHandler) DecodeAndValidate(r *http.Request, model interface{}) error {
	err := me.Decode(r, model)
	if err != nil {
		return err
	}
	modelValidator := &validator.ModelValidator{}
	appErr := modelValidator.Validate(model)
	if appErr != nil {
		return appErr
	}
	return nil
}

//URLParam decode URL param by specified key
func (me *RequestHandler) URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
