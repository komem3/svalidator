package svalidator

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidType     = fmt.Errorf("input is unexpected type")
	ErrNotExistsField  = fmt.Errorf("field does not exist")
	ErrNotEqual        = fmt.Errorf("input value is not equal expected value")
	ErrEmpty           = fmt.Errorf("input value is required")
	ErrTooBig          = fmt.Errorf("input value is too big")
	ErrTooSmall        = fmt.Errorf("input value is too small")
	ErrMismatchPattern = fmt.Errorf("input value is mismatch expected pattern")
)

// ErrValidate is returned on validation error.
// Each validate error is wrapped by this.
type ErrValidate struct {
	Err   error
	Input any
}

func (e *ErrValidate) Error() string {
	return e.Err.Error()
}

func (e *ErrValidate) Is(err error) bool {
	return errors.Is(e.Err, err)
}

func (e *ErrValidate) Unwrap() error {
	return e.Err
}

// ErrObject is returned on object validation error.
type ErrObject []*ErrObjectField

// ErrObjectField contains error and the name of the field in which the error occurred.
type ErrObjectField struct {
	Field string
	Err   error
}

func newErrObject(errs ...*ErrObjectField) error {
	if len(errs) == 0 {
		return nil
	}
	return ErrObject(errs)
}

func newErrObjectField(field string, err error) *ErrObjectField {
	return &ErrObjectField{Field: field, Err: err}
}

func (e ErrObject) Error() string {
	var str strings.Builder
	for i, err := range e {
		fmt.Fprintf(&str, "%s field: %v", err.Field, err.Err)
		if i != len(e)-1 {
			str.WriteRune('\n')
		}
	}
	return str.String()
}

func (e ErrObject) Is(err error) bool {
	errs, ok := err.(ErrObject)
	if !ok {
		return false
	}
	for i, errObj := range errs {
		if len(e)-1 < i {
			return false
		}
		indexErr := e[i]
		if indexErr.Field != errObj.Field || !errors.Is(indexErr.Err, errObj.Err) {
			return false
		}
	}
	return true
}
