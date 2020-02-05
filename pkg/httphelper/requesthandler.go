package httphelper

import (
	"encoding/json"
	"net/http"
)

//RequestDecoder decode http request to a spesific model
type RequestDecoder interface {
	Decode(r *http.Request, decodeTo interface{}) error
}

//RequestHandler implements RequestDecoder
type RequestHandler struct {
}

//Decode decode request body into specified param
func (me *RequestHandler) Decode(r *http.Request, decodeTo interface{}) error {
	err := json.NewDecoder(r.Body).Decode(decodeTo)
	return err
}
