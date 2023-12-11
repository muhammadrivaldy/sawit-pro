package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		if isPublicEndpoint(ctx) {
			return next(ctx)
		}

		// get value authorization from header
		authorization := ctx.Request().Header.Get("authorization")
		if ok := strings.Contains(authorization, "Bearer "); !ok {
			err := errors.New(strings.ToLower(http.StatusText(http.StatusUnauthorized)))
			payloads.ResponseError(ctx, http.StatusUnauthorized, err, nil)
			return err
		}

		// split value without bearer
		authorization = strings.Split(authorization, "Bearer ")[1]

		userId, err := ParseJWT(authorization)
		if err != nil {
			err := errors.New(strings.ToLower(http.StatusText(http.StatusUnauthorized)))
			payloads.ResponseError(ctx, http.StatusUnauthorized, err, nil)
			return err
		}

		ctx.Set("user-id", userId)

		return next(ctx)

	}

}

// isPublicEndpoint is a function to validate the endpoint doesn't need token for authorization
func isPublicEndpoint(ctx echo.Context) bool {

	method := ctx.Request().Method
	path := ctx.Request().URL.Path

	if method == http.MethodPost && path == "/users/login" {
		return true
	} else if method == http.MethodPost && path == "/users" {
		return true
	}

	return false

}
