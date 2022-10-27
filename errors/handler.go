package errors

import (
	"context"
	"encoding/json"
	"net/http"
)

// Error handler for go-kit, specified in the settings when starting the server
// kithttp.ServerErrorEncoder(errors.ErrorEncoder)

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {

	var errorEnc IError
	defaultError := New(TError{HttpStatusCode: 400, ErrCode: "SystemError", ErrMessage: "Something wrong"})

	//Проводим правильное преобразование типа
	//ссылки либо значения
	switch err.(type) {
	case nil:
		errorEnc = defaultError.SetDevMessage("Error is nil")
	case TError:
		errorEnc = err.(TError)
	case *TError:
		errorEnc = *(err.(*TError))
	default:
		errorEnc = defaultError.SetDevMessage(err.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errorEnc.GetHttpStatusCode())
	json.NewEncoder(w).Encode(errorEnc)
}
