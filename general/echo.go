package general

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	code = http.StatusInternalServerError
	msg  string
)

// EchoRestfulErrorHandler - error handler
func EchoRestfulErrorHandler(err error, c echo.Context) {
	log.Error(err)

	if resp, ok := err.(*ErrorResp); ok {
		code = resp.Code
		msg = resp.Message
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				log.Error(err)
			}
		} else {
			err := c.JSON(code, NewErrorWithMessage(code, msg))

			if err != nil {
				log.Error(err)
			}
		}
	}
}
