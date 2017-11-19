package model

import (
	"testing"

	"github.com/Jsharkc/RedPacket/orm"
	"time"
)

func TestInit(t *testing.T) {
	orm.InitTestOrm()

	TimerWorkService.Init(20, 20, 2)
	time.Sleep(time.Second*2)
	t.Fail()
}