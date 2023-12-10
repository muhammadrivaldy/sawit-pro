package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// (GET /users/:id)
func (s *Server) GetUsersId(ctx echo.Context, id int) error {

	user, err := s.repo.SelectUsersById(ctx.Request().Context(), id)
	if err == gorm.ErrRecordNotFound {
		err := errors.New(strings.ToLower(http.StatusText(http.StatusForbidden)))
		payloads.ResponseError(ctx, http.StatusForbidden, err, nil)
		return err
	} else if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	}

	payloads.ResponseOK(ctx, http.StatusOK, payloads.ResponseGetUsersId{
		Id:          user.Id,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	})

	return nil

}
