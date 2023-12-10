package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// (PUT /users/:id)
func (s *Server) PutUsersId(ctx echo.Context, id int) error {

	userId := ctx.Get("user-id").(int)

	if id != userId {
		err := errors.New(strings.ToLower(http.StatusText(http.StatusForbidden)))
		payloads.ResponseError(ctx, http.StatusForbidden, err, nil)
		return err
	}

	param := payloads.RequestPutUsersId{
		PhoneNumber: ctx.FormValue("phone_number"),
		FullName:    ctx.FormValue("full_name"),
	}

	validator := utils.NewValidation()

	err := validator.ValidationStruct(param)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusBadRequest, err, nil)
		return err
	}

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

	userByPhoneNumber, err := s.repo.SelectUsersByPhoneNumber(ctx.Request().Context(), param.PhoneNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	} else if userByPhoneNumber.Id > 0 && user.Id != userByPhoneNumber.Id {
		err := errors.New(strings.ToLower(http.StatusText(http.StatusConflict)))
		payloads.ResponseError(ctx, http.StatusConflict, err, nil)
		return err
	}

	user.FullName = param.FullName
	user.PhoneNumber = param.PhoneNumber
	user.UpdatedBy = null.NewInt(int64(user.Id), true)
	user.UpdatedAt = null.NewTime(time.Now(), true)

	user, err = s.repo.UpdateUsers(ctx.Request().Context(), user)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	}

	payloads.ResponseOK(ctx, http.StatusOK, nil)

	return nil

}
