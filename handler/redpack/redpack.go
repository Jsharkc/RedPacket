package redpack

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Jsharkc/RedPacket/general"
	"github.com/Jsharkc/RedPacket/middleware"
	"github.com/Jsharkc/RedPacket/model"
)

// Create - red packet
func Create(c echo.Context) error {
	var (
		mCreateRedPack model.CreateRedPack
		mPassword      *string
		mRpID          int64
		mUID           int64
		err            error
	)

	mUID = middleware.UserID(c)
	mCreateRedPack.UserID = mUID

	if err = c.Bind(&mCreateRedPack); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(&mCreateRedPack); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if mCreateRedPack.TotalMoney < uint64(mCreateRedPack.Number) {
		return general.NewErrorWithMessage(http.StatusBadRequest, general.ErrRedPackTooManyPeople.Error())
	}

	mRpID, mPassword, err = model.RedPackService.Create(&mCreateRedPack)
	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{general.RespRPID: mRpID, general.RespPwd: *mPassword})
}

// Grab - grab red packet handler
func Grab(c echo.Context) error {
	var (
		mGrabRedPack model.GrabRedPack
		mMoney       uint64
		mUID         int64
		err          error
	)

	mUID = middleware.UserID(c)
	mGrabRedPack.UserID = mUID

	if err = c.Bind(&mGrabRedPack); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(&mGrabRedPack); err != nil {
		return general.NewErrorWithMessage(http.StatusBadRequest, err.Error())
	}

	mMoney, err = model.RedPackService.Grab(&mGrabRedPack)
	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{general.RespMoney: mMoney})
}

// GrabList - obtain grabed red packet list
func GrabList(c echo.Context) error {
	var (
		mGrabList []model.GrabList
		mUID      int64
		err       error
	)

	mUID = middleware.UserID(c)

	mGrabList, err = model.GrabRecordService.GrabList(mUID)
	if err != nil {
		return general.NewErrorWithMessage(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{general.RespGrabList: mGrabList})
}
