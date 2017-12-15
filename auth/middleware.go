package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func MiddlewareRes(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		fmt.Println(c.Get("user"))
		claims := user.Claims.(*JwtUserClaim)

		claims.ExpiresAt = time.Now().Add(time.Minute * TokenExpiryTime).Unix()

		// Generate encoded token and send it as response.
		t, err := user.SignedString([]byte(JwtSecret))
		if err != nil {
			return err
		}

		// grab user id
		id := c.Param("id")

		if id != claims.ID.Hex() {
			return echo.ErrUnauthorized
		}

		c.Response().Header().Set("x_auth_token", t)

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
