package main

import (
	jwt "github.com/Jsharkc/RedPacket/middleware"
	"github.com/Jsharkc/RedPacket/general"
	"github.com/Jsharkc/RedPacket/router"
	"github.com/Jsharkc/RedPacket/orm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/Jsharkc/RedPacket/model"
)

var (
	server *echo.Echo
)

func startEchoServer() {
	server = echo.New()

	server.Use(middleware.Recover())
	server.Use(jwt.CustomJWT())

	server.HTTPErrorHandler = general.EchoRestfulErrorHandler
	server.Validator = general.NewEchoValidator()

	if conf.isDebug {
		log.SetLevel(log.DEBUG)
	} else {
		log.SetLevel(log.INFO)
	}

	router.InitRouter(server)

	server.Start(conf.address)
}

func init() {
	readConfiguration()
	jwt.InitJWTWithToken(conf.tokenKey)
	orm.InitOrm()
	model.TimerWorkService.Init(conf.timerWorkDispatchChanCache, conf.timerWorkChanCache, conf.timerWorkCount)
	startEchoServer()
}
