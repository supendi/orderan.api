package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/supendi/orderan.api/pkg/validator"
)

//RequestDecoder decode http request to a spesific model
type RequestDecoder interface {
	DecodeBody(r *http.Request, decodeTo interface{}) error
	DecodeBodyAndValidate(r *http.Request, model interface{}) error
	DecodeURLParam(r *http.Request, key string) string
}

//RequestHandler implements RequestDecoder
type RequestHandler struct {
}

//DecodeBody decode request body into specified param
func (me *RequestHandler) DecodeBody(r *http.Request, decodeTo interface{}) error {
	err := json.NewDecoder(r.Body).Decode(decodeTo)
	return err
}

//DecodeBodyAndValidate decode the request body, and then validates the struct
func (me *RequestHandler) DecodeBodyAndValidate(r *http.Request, model interface{}) error {
	err := me.DecodeBody(r, model)
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

//DecodeURLParam decode URL param by specified key
func (me *RequestHandler) DecodeURLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
