package linkerror

import "errors"

var (
	JSONError        = errors.New("json error")
	NumberParseError = errors.New("number parse error")
)
