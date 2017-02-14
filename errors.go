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

var Type = errors.New

func New(errType error, msg string) (*Error) {
	return &Error{Type: errType, Msg: msg}
}
func NewWith(errType error, msg string, err *Error) (*Error) {
	return &Error{Type: errType, Msg: msg, prev: err }
}
