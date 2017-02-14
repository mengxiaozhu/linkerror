package linkerror

import (
	"bytes"
	"errors"
)

type Error struct {
	Type      error
	Msg       string
	prev      *Error
	bufString string
}

func (e *Error) Error() (string) {
	if e.bufString != "" {
		return e.bufString
	}
	if e.prev == nil {
		e.bufString = e.Type.Error() + ":" + e.Msg
		return e.bufString
	}
	buf := bytes.NewBuffer(nil)
	linkedErrors := e
	for linkedErrors != nil {
		buf.WriteString(linkedErrors.Type.Error())
		buf.WriteByte(':')
		buf.WriteString(e.Msg)
		buf.WriteByte('\n')
		linkedErrors = linkedErrors.prev
	}
	e.bufString = buf.String()
	return e.bufString
}
func (e *Error) Extend(errType error, msg string) (*Error) {
	return &Error{Type: errType, Msg: msg, prev: e }
}
func (e *Error) Catch(errs ... error) (bool) {
	if e == nil {
		return false
	}
	for _, err := range errs {
		if e.Type == err {
			return true
		}
	}
	return false
}

var Type = errors.New

func New(errType error, msg string) (*Error) {
	return &Error{Type: errType, Msg: msg}
}
func Extend(errType error, msg string, err *Error) (*Error) {
	return &Error{Type: errType, Msg: msg, prev: err }
}

func NewWith(errType error, msg string, err error) (*Error) {
	return &Error{Type: errType, Msg: msg, prev: &Error{Type: err} }
}
