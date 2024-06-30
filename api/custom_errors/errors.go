package custom_errors

import "net/http"

type Error struct {
	status  int
	Success bool              `json:"success" default:"false"`
	Msg     string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e Error) Error() string {
	return e.Msg
}

func (e Error) Status() int {
	return e.status
}

func NewErr(code int, Msg string) Error {
	return Error{
		status: code,
		Msg:    Msg,
	}
}

func NotFound() Error {
	return Error{
		status: http.StatusNotFound,
		Msg:    "not found",
	}
}

func InvalidCredentials() Error {
	return Error{
		status: http.StatusUnauthorized,
		Msg:    "invalid credentials",
	}
}

func Unauthorized() Error {
	return Error{
		status: http.StatusUnauthorized,
		Msg:    "unauthorized",
	}
}

func Forbidden() Error {
	return Error{
		status: http.StatusForbidden,
		Msg:    "user does not have access to this action",
	}
}

func Validation() Error {
	return Error{
		status: http.StatusUnprocessableEntity,
		Msg:    "the provided payload is not valid",
	}
}

func Internal() Error {
	return Error{
		status: http.StatusInternalServerError,
		Msg:    "some thing went wrong, please try later",
	}
}
