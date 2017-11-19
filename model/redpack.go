package model

import (
	"math/rand"
	"time"

	"github.com/Jsharkc/RedPacket/general"
	"github.com/Jsharkc/RedPacket/orm"
	"github.com/Jsharkc/RedPacket/utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type redPackServiceProvider struct{}

var (
	// RedPackService handles operations on model RedPack.
	RedPackService = &redPackServiceProvider{}
)

// RedPack - red packet struct
type RedPack struct {
	ID         int64     `json:"id"       gorm:"column:id;primary_key"`
	UserID     int64     `json:"uid"      gorm:"column:uid"`
	TotalMoney uint64    `json:"total"    gorm:"column:totalmoney"`
	Number     uint32    `json:"num"      gorm:"column:num"`
	Blessing   string    `json:"blessing" gorm:"column:blessing"`
	Password   string    `json:"pass"     gorm:"column:password"`
	Status     uint8     `json:"status"   gorm:"column:status"`
	Created    time.Time `json:"created"  gorm:"column:created"`
}

// CreateRedPack - create red packet struct
type CreateRedPack struct {
	UserID     int64  `json:"uid"`
	TotalMoney uint64 `json:"total"       validate:"gt=0"`
	Number     uint32 `json:"num"         validate:"gt=0"`
	Blessing   string `json:"blessing"    validate:"required"`
}

// GrabRedPack - grab red packet struct
type GrabRedPack struct {
	ID       int64  `json:"rpid"          validate:"gt=0"`
	UserID   int64  `json:"uid"`
	Password string `json:"pwd"           validate:"len=8"`
}

// TableName - returns table name in database
func (RedPack) TableName() string {
	return "redpack"
}

// Create - create red packet
func (rs *redPackServiceProvider) Create(cRedPack *CreateRedPack) (int64, *string, error) {
	var pwd = utils.GenePwd()
	mRedPack := RedPack{
		UserID:     cRedPack.UserID,
		TotalMoney: cRedPack.TotalMoney,
		Number:     cRedPack.Number,
		Blessing:   cRedPack.Blessing,
		Password:   pwd,
		Status:     general.RPGrab,
		Created:    time.Now(),
	}

	err := orm.DBConn.Create(&mRedPack).Error
	if err == nil {
		TimerWorkDispatchChan <- &mRedPack
	}

	return mRedPack.ID, &pwd, err
}

// Grab - grab red packet
func (rs *redPackServiceProvider) Grab(gRedPack *GrabRedPack) (uint64, error) {
	var (
		mAverage    uint64
		mDeserve    uint64
		mRedPack    RedPack
		mGrabRecord = GrabRecord{UserID: gRedPack.UserID, RedPackID: gRedPack.ID}
		err         error
	)
	tx := orm.DBConn.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	err = tx.Where("uid = ? AND rpid = ?", gRedPack.UserID, gRedPack.ID).First(&mGrabRecord).Error

	if err == nil {
		err = general.ErrGrabbed

		return general.IntZero, err
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(err)
		return general.IntZero, err
	}

	err = tx.Raw("SELECT totalmoney, num, status, password FROM redpack WHERE id = ? FOR UPDATE", gRedPack.ID).Scan(&mRedPack).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = general.ErrRedPackNotExist
		}

		log.Error(err)
		return general.IntZero, err
	}

	if mRedPack.Password != gRedPack.Password {
		err = general.ErrRedPackPassword

		log.Error(err)
		return general.IntZero, err
	}

	if mRedPack.Status != general.RPGrab {
		return general.IntZero, general.ErrRedPackFinish
	}

	if mRedPack.Number == general.IntOne {
		mDeserve = mRedPack.TotalMoney

		err = tx.Model(RedPack{ID: gRedPack.ID}).Updates(map[string]interface{}{general.TBTotalMoney: gorm.Expr(general.TBTotalMoneyExprSub, mDeserve),
			general.TBNumber: gorm.Expr(general.TBNumberExprSub, general.IntOne), general.TBStatus: general.RPFinish}).Error
	} else {
		rand.Seed(time.Now().UnixNano())

		mAverage = mRedPack.TotalMoney / uint64(mRedPack.Number)
		mDeserve = uint64(rand.Int63n(2*int64(mAverage)-1) + 1)

		if mRedPack.TotalMoney-mDeserve < uint64(mRedPack.Number-1) {
			mDeserve = mRedPack.TotalMoney - uint64(mRedPack.Number) + 1
		}

		err = tx.Model(RedPack{ID: gRedPack.ID}).Updates(map[string]interface{}{general.TBTotalMoney: gorm.Expr(general.TBTotalMoneyExprSub, mDeserve),
			general.TBNumber: gorm.Expr(general.TBNumberExprSub, general.IntOne)}).Error
	}

	if err != nil {
		log.Error(err)
		return general.IntZero, err
	}

	err = tx.Model(User{UserID: gRedPack.UserID}).Updates(map[string]interface{}{general.TBBalance: gorm.Expr(general.TBBalanceExprAdd, mDeserve)}).Error
	if err != nil {
		log.Error(err)
		return general.IntZero, err
	}

	mGrabRecord.Money = mDeserve
	err = tx.Create(&mGrabRecord).Error

	return mDeserve, err
}

// getActiveRedPack - obtain red packet which status is active
func getActiveRedPack() ([]RedPack, error) {
	var (
		mRedPackList []RedPack
		err          error
	)

	err = orm.DBConn.Where("status = ?", general.RPGrab).Find(&mRedPackList).Error

	return mRedPackList, err
}
