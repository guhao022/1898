package web

import (
	"encoding/json"
	"net/http"
)

type ERR_CODE int

const (
	CODE_SUCCESS ERR_CODE = iota + 2000 // 成功

	CODE_FAIL ERR_CODE = iota + 4000 // 失败
)

type JSON struct {
	ID     string      `json:"id"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	ErrCode int         `json:"errcode"`
}

func Push(w http.ResponseWriter, msg string, data interface{}) {

	j := &JSON{"ok", msg, data, 0}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(j)

}

type Error struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`

	ErrCode int   `json:"errcode"`
}

func (err Error) Error() string {
	return err.Detail
}

func Errors(w http.ResponseWriter, err *Error) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(err)
}

var (
	ErrBadRequest = &Error{"bad_request", "Bad request", "Request body is not well-formed. It must be JSON.", 400}
	ErrNoData     = &Error{"no_data", "No data", `Key "data" in the top level of the JSON document is missing or contains no data`, 422}
	ErrNotFound   = &Error{"not_found", "Not found", "Route not found.", 404}
)

func ErrMissParam(param string, errcode int) *Error {
	return &Error{
		"miss_param",
		"Miss Param",
		"miss query param: " + param,
		errcode,
	}
}

func ErrForbidden(detail string, errcode int) *Error {
	return &Error{
		"forbidden",
		"Forbidden",
		"Forbidden: " + detail,
		errcode,
	}
}

func ErrInternalServer(detail string, errcode int) *Error {
	return &Error{
		"internal_server_error",
		"Internal Server Error",
		detail,
		errcode,
	}
}

func ErrUnauthorized(detail string, errcode int) *Error {
	return &Error{
		"unauthorized",
		"Unauthorized",
		detail,
		errcode,
	}
}
