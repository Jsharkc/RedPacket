package user

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Jsharkc/RedPacket/general"
	"github.com/Jsharkc/RedPacket/middleware"
	"github.com/Jsharkc/RedPacket/model"
)

// GenerToken - Generate jwt token
func GenerToken(c echo.Context) error {
	var (
		mUser model.User
		err   error
	)

	if err = c.Bind(&mUser); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(&mUser); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	token, err := model.UserService.SaveUser(mUser.UserID)
	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{general.RespTokenKey: token})
}

// Balance - obtain user balance
func Balance(c echo.Context) error {
	var (
		mBalance uint64
		mUID     int64
		err      error
	)

	mUID = middleware.UserID(c)

	mBalance, err = model.UserService.Balance(mUID)
	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{general.RespBalance: mBalance})
}
