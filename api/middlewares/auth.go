package middlewares

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"strings"
)

func verifyFirebaseIDToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) {
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	return auth.VerifyIDToken(context.Background(), token)
}

func FirebaseGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, err := verifyFirebaseIDToken(c, authClient)
			if err != nil {
				logrus.Error(err)
				return c.JSON(fasthttp.StatusUnauthorized, "Not Authenticated")
			}
			c.Set("auth", jwtToken)
			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}

func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, _ := verifyFirebaseIDToken(c, authClient)
			c.Set("auth", jwtToken)
			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}