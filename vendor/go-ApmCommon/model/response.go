package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Errors
var ErrUserNotExist = errors.New("user not exist")

// Status Code
const (
	HttpStatusAccepted            = 202
	HttpStatusInternalServerError = 500
)

type ID string

type Response struct {
	Id    ID
	User  *User
	Error *ResponseError
}

func (r ResponseError) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, r.Err)), nil
}

func (r *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v == nil {
		r.Err = nil
		return nil
	}
	switch p := v.(type) {
	case string:
		if r.Err == ErrUserNotExist {
			r.Err = ErrUserNotExist
			return nil
		}
		r.Err = errors.New(p)
		return nil
	default:
		return errors.New("unexpected response error")
	}
	return nil
}

type ResponseError struct {
	Err  error
	Code int
}

func (r ResponseError) Error() string {
	return r.Err.Error()
}

func NewResponseError(e error, c int) *ResponseError {
	return &ResponseError{e, c}
}
