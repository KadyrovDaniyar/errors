package errors

import (
	"net/http"
)

const (
	Success string = "success"
	Failure string = "failure"
)

type IError interface {
	Error() string
	Set(errMessage string) IError
	SetDevMessage(devMessage string) IError
	GetHttpStatusCode() int
}

type TError struct {
	ErrMessage       string            `json:"errMessage" example:"don't have permission to access this resource"`
	DevMessage       string            `json:"devMessage" example:"don't have permission to access colvir user test"`
	ErrCode          string            `json:"errCode" example:"forbidden"`
	Result           string            `json:"result" example:"failure"`
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
	HttpStatusCode   int               `json:"-"`
}

type ValidationError struct {
	Code    string `json:"code" example:"required"`
	Field   string `json:"field" example:"payer.account.number"`
	Message string `json:"message" example:"field is empty"`
}

var (
	BadRequest   = TError{HttpStatusCode: http.StatusBadRequest, ErrCode: "BadRequest"}
	Unauthorized = TError{HttpStatusCode: http.StatusUnauthorized, ErrCode: "Unauthorized"}
	TokenExpired = TError{HttpStatusCode: http.StatusUnauthorized, ErrCode: "TokenExpired"}
	Forbidden    = TError{HttpStatusCode: http.StatusForbidden, ErrCode: "Forbidden"}
	NotFound     = TError{HttpStatusCode: http.StatusNotFound, ErrCode: "NotFound"}
)

func (e TError) Error() string {
	return e.ErrMessage
}

func (e TError) GetHttpStatusCode() int {
	return e.HttpStatusCode
}

func (e TError) Set(errMessage string) IError {
	e.ErrMessage = errMessage
	return e
}

func (e TError) SetDevMessage(devMessage string) IError {
	e.DevMessage = devMessage
	return e
}

func New(e TError) IError {
	e.Result = Failure
	return &e
}
