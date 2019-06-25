package response

import (
	"encoding/json"
	"errors"
	"fmt"

	"go-ApmCommon/model"
)

var ErrUserNotExist = errors.New("user not exist")

const (
	HttpStatusAccepted            = 202
	HttpStatusInternalSErverError = 500
)

type ID string

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

type Response struct {
	Id    ID
	User  *model.User
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
