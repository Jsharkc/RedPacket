package model

import (
	"github.com/Jsharkc/RedPacket/orm"
)

type grabRecordServiceProvider struct{}

var (
	// GrabRecordService handles operations on model GrabRecord.
	GrabRecordService = &grabRecordServiceProvider{}
)

// GrabRecord - grab red packet record struct
type GrabRecord struct {
	ID        int64  `json:"id"       gorm:"column:id;primary_key"`
	UserID    int64  `json:"uid"      gorm:"column:uid"`
	RedPackID int64  `json:"rpid"     gorm:"column:rpid"`
	Money     uint64 `json:"money"    gorm:"column:money"`
}

// GrabList - grab red packet record list struct
type GrabList struct {
	UserID    int64  `json:"uid"      gorm:"column:uid"`
	RedPackID int64  `json:"rpid"     gorm:"column:rpid"`
	Blessing  string `json:"blessing" gorm:"column:blessing"`
	Money     uint64 `json:"money"    gorm:"column:money"`
}

// TableName - returns table name in database
func (GrabRecord) TableName() string {
	return "grabrecord"
}

// GrabList - obtain grab red packet list
func (gs *grabRecordServiceProvider) GrabList(uid int64) ([]GrabList, error) {
	var (
		mGrabList []GrabList
		err       error
	)

	raw := orm.DBConn.Raw("SELECT gr.rpid, gr.money, rp.uid, rp.blessing FROM grabrecord AS gr, redpack AS rp WHERE gr.uid = ? AND gr.rpid = rp.id", uid)
	err = raw.Scan(&mGrabList).Error

	return mGrabList, err
}
