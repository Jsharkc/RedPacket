package model

import (
	"testing"

	"github.com/Jsharkc/RedPacket/orm"
)

func TestGrabList(t *testing.T) {
	orm.InitTestOrm()

	_, err := GrabRecordService.GrabList(12)

	if err != nil {
		t.Error("Obtain grab list error:", err)
	}
}
