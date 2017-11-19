package general

import (
	"errors"
)

var (
	ErrGrabbed              = errors.New("You have grabed red packet!")
	ErrRedPackNotExist      = errors.New("Red packet don't exist!")
	ErrRedPackPassword      = errors.New("Red packet password error!")
	ErrRedPackFinish        = errors.New("Red packet is gone!")
	ErrRedPackTooManyPeople = errors.New("Red packet too many people!")
)
