package httphelper

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/supendi/orderan.api/pkg/errors"
)

//ResponseWriter response writer
type ResponseWriter interface {
	Write(httpSuccessStatus int, result interface{}, err error, w http.ResponseWriter)
}

//ResponseHandler implements ResponseWriter
type ResponseHandler struct {
}

func (me *ResponseHandler) Write(httpSuccessStatus int, result interface{}, err error, w http.ResponseWriter) {
	if err != nil {
		ErrorResponse(err, w)
		return
	}
	WriteResponse(httpSuccessStatus, result, w)
}

//WriteResponse write json response
func WriteResponse(httpStatus int, v interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if v == nil || reflect.ValueOf(v).IsNil() {
		w.WriteHeader(httpStatus)
		w.Write([]byte{})
		return
	}

	jsonBody, err := json.Marshal(v)

	if err != nil {
		ErrorResponse(err, w)
		return
	}

	w.WriteHeader(httpStatus)
	w.Write(jsonBody)

}

//ErrorResponse write error response
func ErrorResponse(err error, w http.ResponseWriter) {
	if err != nil {
		if errors.IsAppError(err) {
			WriteResponse(400, err, w)
			return
		}
		WriteResponse(500, err, w)
	}
}
