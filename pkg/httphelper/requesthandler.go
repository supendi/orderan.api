package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

//RequestDecoder decode http request to a spesific model
type RequestDecoder interface {
	Decode(r *http.Request, decodeTo interface{}) error
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

//URLParam decode URL param by specified key
func (me *RequestHandler) URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
