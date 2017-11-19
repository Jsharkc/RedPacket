package model

import (
	"github.com/Jsharkc/RedPacket/middleware"
	"github.com/Jsharkc/RedPacket/orm"
)

type userServiceProvider struct{}

var (
	// UserService handles operations on model User.
	UserService = &userServiceProvider{}
)

// User - user struct
type User struct {
	UserID  int64  `json:"id" gorm:"column:id;primary_key" validate:"required"`
	Balance uint64 `json:"balance"`
}

// TableName - returns table name in database
func (User) TableName() string {
	return "user"
}

// SaveUser create a new user wallet
func (us *userServiceProvider) SaveUser(uid int64) (string, error) {
	var mUser = User{UserID: uid}

	token, err := middleware.NewToken(uid)
	if err != nil {
		return token, err
	}

	return token, orm.DBConn.FirstOrCreate(&mUser).Error
}

// Balance - obtain user balance
func (us *userServiceProvider) Balance(uid int64) (uint64, error) {
	var (
		mUser User
		err   error
	)

	err = orm.DBConn.Where("id = ?", uid).First(&mUser).Error

	return mUser.Balance, err
}
