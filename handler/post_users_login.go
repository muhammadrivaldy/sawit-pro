package handler

import (
	"errors"
	"net/http"

	"github.com/SawitProRecruitment/UserService/payloads"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// (POST /users/login)
func (s *Server) PostUsersLogin(ctx echo.Context) error {

	phoneNumber, password, _ := ctx.Request().BasicAuth()

	param := payloads.RequestPostUsersLogin{
		PhoneNumber: phoneNumber,
		Password:    password,
	}

	validator := utils.NewValidation()

	err := validator.ValidationStruct(param)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusBadRequest, err, nil)
		return err
	}

	user, err := s.repo.SelectUsersByPhoneNumber(ctx.Request().Context(), param.PhoneNumber)
	if err == gorm.ErrRecordNotFound {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusBadRequest, errors.New("your email & password are wrong"), nil)
		return err
	} else if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	}

	err = utils.ComparePassword(user.PasswordHash, param.Password)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusBadRequest, errors.New("your email & password are wrong"), nil)
		return err
	}

	token, err := utils.CreateJWT(user.Id)
	if err != nil {
		ctx.Logger().Error(err)
		payloads.ResponseError(ctx, http.StatusInternalServerError, err, nil)
		return err
	}

	payloads.ResponseOK(ctx, http.StatusOK, payloads.ResponsePostUsersLogin{
		Token: token,
	})

	return nil

}
