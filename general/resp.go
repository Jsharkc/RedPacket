package general

import (
	"net/http"
)

type ErrorResp struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func NewErrorWithMessage(code int, msg string) *ErrorResp {
	if code == http.StatusOK {
		msg = ""
	}

	return &ErrorResp{
		Code:    code,
		Message: msg,
	}
}

func (this *ErrorResp) Error() string {
	return this.Message
}
