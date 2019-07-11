package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Errors
var ErrUserNotExist = errors.New("user not exist")

// Status Code
//const (
//	HttpStatusAccepted            = 202
//	HttpStatusInternalServerError = 500
//)

//TODO use defined error
var (
	Accepted            = ResponseError{nil, 202}
	InternalServerError = ResponseError{ErrUserNotExist, 500}
)
//
//
type ResponseError struct {
	Err  error
	Code int
}

func NewResponseError(e error, c int) *ResponseError {
		return &ResponseError{e, c}
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


func (r ResponseError) Error() string {
		return fmt.Sprintf(`{"message":"%s", "code":"%d"}`, r.Err, r.Code)
}

//
//
type ID string

type Response struct {
	Id    ID
	User  *User
	Error *ResponseError
}
