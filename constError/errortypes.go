// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package constError

import (
	"errors"
	stderror "errors"
	"fmt"
	"strings"
)

type ConstError struct {
	Message string
	Code    int
}

// ConstError implements error
func (e ConstError) Error() string {
	return e.Message
}

// Different types of errors
var (
	Timeout = ConstError{Message: "Timeout", Code: 1001}
)

// errWithType is an Err bundled with its error type (a ConstError)
type errWithType struct {
	error
	errType ConstError
}

// Is compares `target` with e's error type
func (e *errWithType) Is(target error) bool {
	if &e.errType == nil {
		return false
	}
	return target == e.errType
}

// Unwrap an errWithType gives the underlying Err
func (e *errWithType) Unwrap() error {
	return e.error
}

func wrapErrorWithMsg(err error, msg string) error {
	if err == nil {
		return stderror.New(msg)
	}
	if msg == "" {
		return err
	}
	return fmt.Errorf("%s: %w", msg, err)
}

func makeWrappedConstError(err error, format string, args ...interface{}) error {
	separator := " "
	if err.Error() == "" || errors.Is(err, &fmtNoop{}) {
		separator = ""
	}
	return fmt.Errorf(strings.Join([]string{format, "%w"}, separator), append(args, err)...)
}

// WithType is responsible for annotating an already existing error so that it
// also satisfies that of a ConstError. The resultant error returned should
// satisfy Is(err, errType). If err is nil then a nil error will also be returned.
//
// Now with Go's Is, As and Unwrap support it no longer makes sense to Wrap()
// 2 errors as both of those errors could be chains of errors in their own right.
// WithType aims to solve some of the usefulness of Wrap with the ability to
// make a pre-existing error also satisfy a ConstError type.
func WithType(err error, errType ConstError) error {
	if err == nil {
		return nil
	}
	return &errWithType{
		error:   err,
		errType: errType,
	}
}

// NewTimeout returns an error which wraps err and satisfies Is(err, Timeout)
// and the Locationer interface.
func NewTimeout(err error, msg string) error {
	return &errWithType{
		error:   newLocationError(wrapErrorWithMsg(err, msg), 1),
		errType: Timeout,
	}
}
