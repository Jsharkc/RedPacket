package router

import (
	"github.com/labstack/echo"

	"github.com/Jsharkc/RedPacket/handler/user"
	"github.com/Jsharkc/RedPacket/handler/redpack"
	"github.com/Jsharkc/RedPacket/middleware"
)

// InitRouter initialize routes.
func InitRouter(server *echo.Echo) {
	if server == nil {
		panic("[InitRouter], server couldn't be nil")
	}

	// User
	server.POST("/gene/token", user.GenerToken)

	group := server.Group("", middleware.MustLoginIn)

	group.GET("/user/balance", user.Balance)

	// Red packet
	group.POST("/redpack/send", redpack.Create)
	group.POST("/redpack/grab", redpack.Grab)
	group.GET("/redpack/grab/list", redpack.GrabList)
}
