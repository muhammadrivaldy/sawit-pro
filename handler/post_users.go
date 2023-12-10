package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// (POST /users)
func (s *Server) PostUsers(ctx echo.Context) error {

	param := payloads.RequestPostUsers{
		PhoneNumber: ctx.FormValue("phone_number"),
		FullName:    ctx.FormValue("full_name"),
		Password:    ctx.FormValue("password"),
	}

	validator := utils.NewValidation()

	err := validator.ValidationStruct(param)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusBadRequest, err, nil)
		return err
	}

	user, err := s.repo.SelectUsersByPhoneNumber(ctx.Request().Context(), param.PhoneNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	} else if user.Id > 0 {
		err := errors.New(strings.ToLower(http.StatusText(http.StatusConflict)))
		payloads.ResponseError(ctx, http.StatusConflict, err, nil)
		return err
	}

	passwordHash, err := utils.GeneratePassword(param.Password)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	}

	user, err = s.repo.InsertUsers(ctx.Request().Context(), models.User{
		FullName:     param.FullName,
		PhoneNumber:  param.PhoneNumber,
		PasswordHash: passwordHash,
		CreatedBy:    0,
		CreatedAt:    time.Now(),
	})
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	}

	payloads.ResponseOK(ctx, http.StatusCreated, payloads.ResponsePostUsers{
		Id:          user.Id,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	})

	return nil

}
