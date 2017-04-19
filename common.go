package linkerror

import "errors"

var (
	JSONError            = errors.New("json error")
	HTTPRequestSendError = errors.New("HTTP request send error")
	NumberParseError     = errors.New("number parse error")
)
