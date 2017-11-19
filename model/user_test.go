package model

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Jsharkc/RedPacket/orm"
)

func TestSaveUser(t *testing.T) {
	orm.InitTestOrm()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	token, err := UserService.SaveUser(r.Int63n(1000))
	t.Logf("Token:%s", token)

	if err != nil {
		t.Error("Save user error!")
	}
}

func TestBalance(t *testing.T) {
	orm.InitTestOrm()

	_, err := UserService.Balance(12)
	if err != nil {
		t.Error("Obtain user balance error!")
	}
}
