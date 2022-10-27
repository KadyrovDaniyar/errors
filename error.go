package errors

type IError interface {
	Error() string
	Set(errMessage string) IError
	SetDevMessage(devMessage string) IError
	GetHttpStatusCode() int
}
