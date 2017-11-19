package middleware

import (
	"time"

	"github.com/Jsharkc/RedPacket/general"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// TokenHMACKey .
	TokenHMACKey string

	urlMap map[string]struct{}
)

// CustomJWT - jwt middleware
func CustomJWT() echo.MiddlewareFunc {
	jwtconf := middleware.JWTConfig{
		Skipper:    CustomSkipper,
		SigningKey: []byte(TokenHMACKey),
	}

	return middleware.JWTWithConfig(jwtconf)
}

// CustomSkipper - custom skipper
func CustomSkipper(c echo.Context) bool {
	if _, ok := urlMap[c.Request().RequestURI]; ok {
		return true
	}

	return false
}

// InitJWTWithToken initialize the HMAC token.
func InitJWTWithToken(token string) {
	TokenHMACKey = token
}

// NewToken generates a JWT token.
func NewToken(uid int64) (string, error) {
	token := jwtgo.New(jwtgo.SigningMethodHS256)

	claims := token.Claims.(jwtgo.MapClaims)
	claims[general.ClaimUID] = uid
	claims[general.ClaimExpire] = time.Now().Add(time.Hour * general.TokenExpireInHour).Unix()

	t, err := token.SignedString([]byte(TokenHMACKey))
	return t, err
}

// UserID returns user identity.
func UserID(c echo.Context) int64 {
	return c.Get(general.ClaimUID).(int64)
}

// MustLoginIn - obtain user id
func MustLoginIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwtgo.Token)
		claims := user.Claims.(jwtgo.MapClaims)
		uid := claims[general.ClaimUID].(float64)

		c.Set(general.ClaimUID, int64(uid))

		return next(c)
	}
}

func init() {
	urlMap = make(map[string]struct{})

	urlMap["/gene/token"] = struct{}{}
}
