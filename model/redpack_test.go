package model

import (
	"testing"

	"github.com/Jsharkc/RedPacket/orm"
	"github.com/labstack/gommon/log"
)

func TestCreate(t *testing.T) {
	orm.InitTestOrm()

	cRP := CreateRedPack{
		UserID:     12,
		TotalMoney: 200,
		Number:     4,
		Blessing:   "Happy birthday to you!",
	}

	rpid, pwd, err := RedPackService.Create(&cRP)
	t.Logf("Redpack ID:%d -- Pwd:%s", rpid, *pwd)

	if err != nil {
		t.Error("Create redpack error:", err)
	}
}

func TestGrab(t *testing.T) {
	orm.InitTestOrm()

	mGRP := GrabRedPack{
		ID:       10030,
		UserID:   13,
		Password: "jkauKX02",
	}

	_, err := RedPackService.Grab(&mGRP)
	if err != nil {
		log.Error(err)
		t.Error("Grab red packet error:", err)
	}
}
