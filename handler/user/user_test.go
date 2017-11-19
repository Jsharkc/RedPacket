package user

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Jsharkc/RedPacket/general"
	"github.com/Jsharkc/RedPacket/model"
	"github.com/Jsharkc/RedPacket/orm"
	"github.com/labstack/echo"
)

func TestGenerToken(t *testing.T) {
	orm.InitTestOrm()

	usr := model.User{UserID: 123}
	server := echo.New()
	server.Validator = general.NewEchoValidator()

	reqBody, _ := json.Marshal(&usr)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/gene/token", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	context := server.NewContext(req, rr)

	err := GenerToken(context)

	if err != nil {
		t.Errorf("Generate token error:%+v", err)
	}
}
